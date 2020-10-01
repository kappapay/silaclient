package sila_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_Accounts(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		userHandle := testConfig.UserHandle
		userWalletPrivateKey := testConfig.UserWalletPrivateKeyHex
		userWalletAddress, err := sila.GetWalletAddress(userWalletPrivateKey)
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.AuthPrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("And the specified integration user exists and has passed KYC", func() {
			response, err := client.CheckHandle(userHandle).Do()
			So(err, ShouldBeNil)
			if response.Success == true {
				registerResponse, err := client.Register(userHandle).
					SetIndividualEntity("Alberta", "Bobbeth", "1950-10-31").
					SetAddress(sila.RegistrationAddress{
						AddressAlias:   "Home",
						StreetAddress1: "1234 Fake St.",
						City:           "Los Angeles",
						State:          "CA",
						Country:        "US",
						PostalCode:     "90001",
					}).
					SetIdentity(sila.Ssn, "181-91-1478").
					SetContact("Home", "123-456-7890", "alberta@bobbeth.com").
					SetCrypto("Main Address", userWalletAddress).
					Do()
				So(err, ShouldBeNil)
				So(registerResponse.Success, ShouldBeTrue)

				requestKycResponse, err := client.RequestKyc(userHandle).Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(requestKycResponse.Success, ShouldBeTrue)

				time.Sleep(30 * time.Second)

				checkKycResponse, err := client.CheckKyc(userHandle).Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(checkKycResponse.Success, ShouldBeTrue)
			}

			Convey("A call to directly link an account should succeed", func() {
				accountName := "My Account"
				response, err := client.LinkAccount(userHandle).
					SetRef("My Reference").
					SetAccountName(accountName).
					SetAccountType("CHECKING").
					SetDirectLinkAccount("123456789012", "123456789").
					Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.Message, ShouldEqual, "Bank account successfully manually linked.")

				Convey("A call to get the Plaid same day auth should fail, since the account was directly linked", func() {
					response, err := client.PlaidSameDayAuth(userHandle, accountName).Do()
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeFalse)
					So(response.Status, ShouldEqual, "FAILURE")
					So(response.ValidationDetails, ShouldBeNil)
				})

				Convey("A call to get the newly linked account's balance should fail since it was directly linked", func() {
					response, err := client.GetAccountBalance(userHandle, accountName).
						SetRef("My Reference").
						Do(userWalletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeFalse)
					So(response.Status, ShouldEqual, "FAILURE")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Reference, ShouldEqual, "My Reference")
					So(response.Message, ShouldEqual, "Cannot fetch account balance for account My Account; account either was not linked with Plaid or has not completed verification.")
				})

				Convey("A call to get the accounts for the user should succeed", func() {
					response, err := client.GetAccounts(userHandle).
						Do(userWalletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Accounts, ShouldNotBeEmpty)
					So(response.Accounts[0].AccountNumber, ShouldEqual, "*9012")
					So(response.Accounts[0].RoutingNumber, ShouldEqual, "123456789")
					So(response.Accounts[0].AccountStatus, ShouldEqual, "active")
					So(response.Accounts[0].Active, ShouldBeTrue)
					So(response.Accounts[0].AccountLinkStatus, ShouldEqual, "unverified_manual_input")
					So(response.Accounts[0].AccountName, ShouldStartWith, accountName)
				})
			})
		})
	})
}

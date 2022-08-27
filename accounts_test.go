package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/kappapay/silaclient"
)

func TestClient_Accounts(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := readTestConfig()
		So(err, ShouldBeNil)
		userHandle := testConfig.UserHandle
		userWalletPrivateKey := testConfig.UserWalletPrivateKeyHex
		userWalletAddress, err := sila.GetWalletAddress(userWalletPrivateKey)
		So(err, ShouldBeNil)
		silaClient, err := sila.NewClient(
			testConfig.AuthPrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("And the specified integration user exists and has passed KYC", func() {
			ensureIntegrationUserExists(silaClient, userHandle, userWalletAddress, userWalletPrivateKey)

			Convey("A call to directly link an account should succeed", func() {
				accountName := "My Account"
				response, err := silaClient.LinkAccount(userHandle).
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
					response, err := silaClient.PlaidSameDayAuth(userHandle, accountName).Do()
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeFalse)
					So(response.Status, ShouldEqual, "FAILURE")
					So(response.ValidationDetails, ShouldBeNil)
				})

				Convey("A call to get the newly linked account's balance should fail since it was directly linked", func() {
					response, err := silaClient.GetAccountBalance(userHandle, accountName).
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
					response, err := silaClient.GetAccounts(userHandle).
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

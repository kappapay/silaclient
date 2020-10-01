package sila_test

import (
	uuid "github.com/satori/go.uuid"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_Wallets(t *testing.T) {
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

			Convey("A call to get a main wallet's balance should succeed", func() {
				response, err := client.GetWalletBalance(userWalletAddress).Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
			})

			Convey("A call to get the main wallet should succeed", func() {
				response, err := client.GetWallet(userHandle).
					SetRef("My Reference").
					Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.IsWhitelisted, ShouldBeTrue)
				So(response.SilaBalance, ShouldBeZeroValue)
				So(response.Wallet.BlockchainAddress, ShouldEqual, userWalletAddress)
			})

			Convey("A call to get a user's wallets should succeed", func() {
				response, err := client.GetWallets(userHandle).
					SetSearchFilters(sila.WalletSearchFilters{
						Page:          1,
						PerPage:       5,
						SortAscending: true,
					}).
					Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Wallets, ShouldNotBeEmpty)
				So(response.Page, ShouldEqual, 1)
				So(response.ReturnedCount, ShouldBeGreaterThanOrEqualTo, 1)
				So(response.TotalCount, ShouldBeGreaterThanOrEqualTo, 1)
				So(response.TotalPageCount, ShouldBeGreaterThanOrEqualTo, 1)
			})

			Convey("A new wallet private key and address are generated for registration, along with a signature", func() {
				newWalletPrivateKey, err := sila.GenerateNewPrivateKey()
				So(err, ShouldBeNil)
				newWalletAddress, err := sila.GetWalletAddress(newWalletPrivateKey)
				So(err, ShouldBeNil)
				signature, err := sila.GenerateWalletSignature([]byte(newWalletAddress), newWalletPrivateKey)

				Convey("A call to register the wallet should succeed", func() {
					response, err := client.RegisterWallet(userHandle).
						SetRef("My Reference").
						SetWallet("My Integration Wallet", newWalletAddress, signature).
						Do(userWalletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					So(response.Status, ShouldEqual, "SUCCESS")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Reference, ShouldEqual, "My Reference")
					So(response.WalletNickname, ShouldStartWith, "My Integration Wallet")

					Convey("A call to update the new wallet's nickname should succeed", func() {
						updatedWalletNickname := uuid.NewV4().String()
						response, err := client.UpdateWallet(userHandle).
							SetRef("My Reference").
							SetNickname(updatedWalletNickname).
							Do(newWalletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.Status, ShouldEqual, "SUCCESS")
						So(response.Message, ShouldEqual, "Wallet updated.")
						So(response.ValidationDetails, ShouldBeNil)
						So(response.Reference, ShouldEqual, "My Reference")
						So(response.Wallet.Nickname, ShouldEqual, updatedWalletNickname)
						So(response.Wallet.Default, ShouldBeFalse)
						So(response.Wallet.BlockchainAddress, ShouldEqual, newWalletAddress)
						So(response.Changes, ShouldNotBeEmpty)

						Convey("A call to delete the new wallet should succeed", func() {
							response, err := client.DeleteWallet(userHandle).
								SetRef("My Reference").
								Do(newWalletPrivateKey)
							So(err, ShouldBeNil)
							So(response.Success, ShouldBeTrue)
							So(response.Status, ShouldEqual, "SUCCESS")
							So(response.ValidationDetails, ShouldBeNil)
							So(response.Reference, ShouldEqual, "My Reference")
						})
					})
				})
			})
		})
	})
}

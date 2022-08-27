package sila_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/kappapay/silaclient"
	"github.com/kappapay/silaclient/domain"
)

func TestClient_Wallets(t *testing.T) {
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

			Convey("A call to get a main wallet's balance should succeed", func() {
				response, err := silaClient.GetWalletBalance(userWalletAddress).Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
			})

			Convey("A call to get the main wallet should succeed", func() {
				response, err := silaClient.GetWallet(userHandle).
					SetRef("My Reference").
					Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.IsWhitelisted, ShouldBeTrue)
				So(response.Wallet.BlockchainAddress, ShouldEqual, userWalletAddress)
			})

			Convey("A call to get a user's wallets should succeed", func() {
				response, err := silaClient.GetWallets(userHandle).
					SetSearchFilters(domain.WalletSearchFilters{
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
					response, err := silaClient.RegisterWallet(userHandle).
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
						response, err := silaClient.UpdateWallet(userHandle).
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
							response, err := silaClient.DeleteWallet(userHandle).
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

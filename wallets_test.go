package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_RegisterWallet(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)

			Convey("The new wallet private key and address are generated, along with a signature", func() {
				newWalletPrivateKey, err := sila.GenerateNewPrivateKey()
				So(err, ShouldBeNil)
				address, err := sila.GetWalletAddress(newWalletPrivateKey)
				So(err, ShouldBeNil)
				signature, err := sila.GenerateWalletSignature([]byte(address), newWalletPrivateKey)

				Convey("The call to register a wallet should succeed", func() {
					response, err := client.RegisterWallet("user.silamoney.eth").
						SetRef("My Reference").
						SetWallet("My Wallet", address, signature).
						Do(privateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					So(response.Status, ShouldEqual, "SUCCESS")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Reference, ShouldEqual, "My Reference")
				})
			})
		})
	})
}

func TestClient_GetWallet(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to get a wallet should succeed", func() {
				response, err := client.GetWallet("user.silamoney.eth").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
			})
		})
	})
}

func TestClient_GetWallets(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("An 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to get wallets should succeed", func() {
				response, err := client.GetWallets("user.silamoney.eth").
					SetRef("My Reference").
					SetSearchFilters(sila.WalletSearchFilters{
						Page:          1,
						PerPage:       5,
						SortAscending: true,
					}).
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
			})
		})
	})
}

func TestClient_UpdateWallet(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to update a wallet should succeed", func() {
				response, err := client.UpdateWallet("user.silamoney.eth").
					SetRef("My Reference").
					SetNickname("Favorite Wallet").
					SetDefault(true).
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
			})
		})
	})
}

func TestClient_GetWalletBalance(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			address, err := sila.GetWalletAddress(privateKey)
			So(err, ShouldBeNil)

			Convey("The call to get a wallet's balance should succeed", func() {
				response, err := client.GetWalletBalance(address).Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
			})
		})
	})
}

func TestClient_DeleteWallet(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to delete a wallet should succeed", func() {
				response, err := client.DeleteWallet("user.silamoney.eth").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
			})
		})
	})
}

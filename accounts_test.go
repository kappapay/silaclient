package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_LinkAccount(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			//address, err := sila.GetWalletAddress(privateKey)
			//So(err, ShouldBeNil)
			Convey("The call to link an account should succeed", func() {
				response, err := client.LinkAccount("user.silamoney.eth").
					SetRef("My Reference").
					SetAccountName("My Account").
					SetAccountType("CHECKING").
					SetDirectLinkAccount("123456789012", "123456789").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_PlaidSameDayAuth(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to plaid same day auth should succeed", func() {
			response, err := client.PlaidSameDayAuth("user.silamoney.eth", "SomeAccountName").Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldEqual, "SUCCESS")
		})
	})
}

func TestClient_GetAccounts(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			//address, err := sila.GetWalletAddress(privateKey)
			//So(err, ShouldBeNil)
			Convey("The call to get accounts should succeed", func() {
				response, err := client.GetAccounts("user.silamoney.eth").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Accounts, ShouldHaveLength)
			})
		})
	})
}

func TestClient_GetAccountBalance(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			//address, err := sila.GetWalletAddress(privateKey)
			//So(err, ShouldBeNil)
			Convey("The call to get an account's balance should succeed", func() {
				response, err := client.GetAccountBalance("user.silamoney.eth", "SomeAccountName").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

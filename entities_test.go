package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

// TODO: These are currently integration tests, and ones that break at that (the auth private key is totally nonsense)

func TestClient_CheckHandle(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to check handle should succeed", func() {
			response, err := client.CheckHandle("user.silamoney.eth").SetRef("My Reference").Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldEqual, "SUCCESS")
		})
	})
}

func TestClient_Register(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			address, err := sila.GetWalletAddress(privateKey)
			So(err, ShouldBeNil)
			Convey("The call to check handle should succeed", func() {
				response, err := client.Register("user.silamoney.eth").
					SetRef("My Reference").
					SetIndividualEntity("Alberta", "Bobbeth", "1950-10-31").
					SetAddress(sila.Address{
						AddressAlias:   "Home",
						StreetAddress1: "1234 Fake St.",
						City:           "Los Angeles",
						State:          "CA",
						Country:        "US",
						PostalCode:     "90001",
					}).
					SetIdentity(sila.SsnIdentity, "181-91-1478").
					SetContact("Home", "123-456-7890", "alberta@bobbeth.com").
					SetCrypto("Main Address", address).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_RequestKyc(t *testing.T) {
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
			Convey("The call to request KYC should succeed", func() {
				response, err := client.RequestKyc("user.silamoney.eth").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_CheckKyc(t *testing.T) {
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
			Convey("The call to check KYC should succeed", func() {
				response, err := client.CheckKyc("user.silamoney.eth").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

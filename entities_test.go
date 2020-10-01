package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_CheckHandle(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to check handle should succeed", func() {
			response, err := client.CheckHandle("ce6ce827-109b-4815-b31a-9af2c3cba031").
				SetRef("My Reference").
				Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Reference, ShouldEqual, "My Reference")
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
		})
	})
}

func TestClient_Register(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			address, err := sila.GetWalletAddress(privateKey)
			So(err, ShouldBeNil)
			Convey("The call to register should succeed", func() {
				response, err := client.Register("dcd3330d.pastel.dev").
					SetRef("My Reference").
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
					SetCrypto("Main Address", address).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
			})
		})
	})
}

func TestClient_RequestKyc(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)

		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to request KYC should succeed", func() {
				response, err := client.RequestKyc("dcd3330d.pastel.dev").
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

func TestClient_CheckKyc(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to check KYC should succeed", func() {
				response, err := client.CheckKyc("dcd3330d.pastel.dev").
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

func TestClient_GetEntity(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The user private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to get an entity should succeed", func() {
				response, err := client.GetEntity("dcd3330d.pastel.dev").
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

func TestClient_GetEntities(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.PrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to get entities should succeed", func() {
			response, err := client.GetEntities().
				SetPage(1).
				SetPerPage(20).
				Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
		})
	})
}

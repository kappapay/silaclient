package sila_test

import (
	uuid "github.com/satori/go.uuid"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_IndividualEntityRegistration(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := readTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.AuthPrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)

		Convey("And a new user handle is chosen, and a wallet private key generated", func() {
			userUuid := uuid.NewV4().String()
			userHandle := userUuid + ".test.golang"

			walletPrivateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			walletAddress, err := sila.GetWalletAddress(walletPrivateKey)
			So(err, ShouldBeNil)

			Convey("A call to check if a known existing handle exists should fail", func() {
				response, err := client.CheckHandle("user.silamoney.eth").
					SetRef("My Reference").
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeFalse)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.Message, ShouldEqual, "user is taken.")
				So(response.Status, ShouldEqual, "FAILURE")
				So(response.ValidationDetails, ShouldBeNil)
			})
			Convey("A call to check if the new user handle is free should succeed", func() {
				response, err := client.CheckHandle(userHandle).
					SetRef("My Reference").
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
			})

			Convey("A call to register the new user should succeed", func() {
				response, err := client.Register(userHandle).
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
					SetCrypto("Main Address", walletAddress).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")

				Convey("A call to request KYC on the newly registered account should succeed", func() {
					response, err := client.RequestKyc(userHandle).
						SetRef("My Reference").
						Do(walletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					So(response.Status, ShouldEqual, "SUCCESS")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Reference, ShouldEqual, "My Reference")

					Convey("A call to check KYC on the just registered account should not succeed", func() {
						response, err := client.CheckKyc(userHandle).
							SetRef("My Reference").
							Do(walletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeFalse)
						So(response.Status, ShouldEqual, "FAILURE")
						So(response.ValidationDetails, ShouldBeNil)
						So(response.Reference, ShouldEqual, "My Reference")
						So(response.EntityType, ShouldEqual, "individual")
						So(response.VerificationStatus, ShouldEqual, "pending")
					})

					Convey("A call to get the newly created entity should succeed", func() {
						response, err := client.GetEntity(userHandle).
							Do(walletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.Status, ShouldEqual, "SUCCESS")
						So(response.ValidationDetails, ShouldBeNil)
						So(response.EntityType, ShouldEqual, "individual")
						So(response.UserHandle, ShouldEqual, userHandle)
						So(response.Entity.FirstName, ShouldEqual, "Alberta")
						So(response.Addresses, ShouldNotBeEmpty)
						So(response.Identities, ShouldNotBeEmpty)
						So(response.Emails, ShouldNotBeEmpty)
						So(response.Phones, ShouldNotBeEmpty)
					})

					Convey("A call to get entities should succeed", func() {
						response, err := client.GetEntities().
							SetPage(1).
							SetPerPage(20).
							Do()
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.Status, ShouldEqual, "SUCCESS")
						So(response.ValidationDetails, ShouldBeNil)
						So(response.Entities.Individuals, ShouldNotBeEmpty)
						So(response.Pagination.ReturnedCount, ShouldBeGreaterThanOrEqualTo, 1)
						So(response.Pagination.TotalCount, ShouldBeGreaterThanOrEqualTo, 1)
						So(response.Pagination.TotalPages, ShouldBeGreaterThanOrEqualTo, 1)
						So(response.Pagination.CurrentPage, ShouldEqual, 1)
					})
				})
			})
		})
	})
}

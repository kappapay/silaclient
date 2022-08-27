package sila_test

import (
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/kappapay/silaclient"
	"github.com/kappapay/silaclient/domain"
)

func TestClient_IndividualEntityRegistration(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := readTestConfig()
		So(err, ShouldBeNil)
		silaClient, err := sila.NewClient(
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
				response, err := silaClient.CheckHandle("user.silamoney.eth").
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
				response, err := silaClient.CheckHandle(userHandle).
					SetRef("My Reference").
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
			})

			Convey("A call to register the new user should succeed", func() {
				response, err := silaClient.Register(userHandle).
					SetRef("My Reference").
					SetIndividualEntity("Alberta", "Bobbeth", "1950-10-31").
					SetAddress(domain.RegistrationAddress{
						AddressAlias:   "Home",
						StreetAddress1: "1234 Fake St.",
						City:           "Los Angeles",
						State:          "CA",
						Country:        "US",
						PostalCode:     "90001",
					}).
					SetIdentity(domain.Ssn, "181-91-1478").
					SetContact("Home", "123-456-7890", "alberta@bobbeth.com").
					SetCrypto("Main Address", walletAddress).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")

				Convey("A call to get the newly created entity should succeed", func() {
					response, err := silaClient.GetEntity(userHandle).
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

					Convey("A call to update registration data should succeed", func() {
						newAddress := domain.RegistrationAddress{
							AddressAlias:   "Home",
							StreetAddress1: "4321 Bogus Ln.",
							City:           "Los Angeles",
							State:          "CA",
							Country:        "US",
							PostalCode:     "90001",
						}
						response, err := silaClient.UpdateRegistrationData(userHandle).
							SetAddress(response.Addresses[0].Uuid, newAddress).
							Do(walletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.Status, ShouldEqual, "SUCCESS")
						So(response.ValidationDetails, ShouldBeNil)
						So(response.Address.Uuid, ShouldNotBeEmpty)
						So(response.Address.StreetAddress1, ShouldEqual, newAddress.StreetAddress1)
						So(response.Address.AddedTime, ShouldNotBeZeroValue)
						So(response.Address.ModifiedTime, ShouldNotBeZeroValue)
					})

					Convey("A call to delete registration data should succeed", func() {
						response, err := silaClient.DeleteRegistrationData(userHandle).
							SetEmail(response.Emails[0].Uuid).
							Do(walletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.Status, ShouldEqual, "SUCCESS")
						So(response.ValidationDetails, ShouldBeNil)
					})

					Convey("A call to add registration data should succeed", func() {
						newEmail := "newEmail@gmail.com"
						response, err := silaClient.AddRegistrationData(userHandle).
							SetEmail(newEmail).
							Do(walletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.Status, ShouldEqual, "SUCCESS")
						So(response.ValidationDetails, ShouldBeNil)
						So(response.Email.Uuid, ShouldNotBeEmpty)
						So(response.Email.Email, ShouldEqual, newEmail)
						So(response.Email.AddedTime, ShouldNotBeZeroValue)
						So(response.Email.ModifiedTime, ShouldNotBeZeroValue)
					})
				})

				Convey("A call to request KYC on the newly registered account should succeed", func() {
					response, err := silaClient.RequestKyc(userHandle).
						SetRef("My Reference").
						Do(walletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					So(response.Status, ShouldEqual, "SUCCESS")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Reference, ShouldEqual, "My Reference")

					Convey("A call to check KYC on the just registered account should not succeed", func() {
						response, err := silaClient.CheckKyc(userHandle).
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

					Convey("A call to get entities should succeed", func() {
						response, err := silaClient.GetEntities().
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

func TestClient_BusinessEntityRegistration(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := readTestConfig()
		So(err, ShouldBeNil)
		silaClient, err := sila.NewClient(
			testConfig.AuthPrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)

		Convey("And the ID's and wallets of the business members are generated", func() {
			adminUuid := uuid.NewV4().String()
			adminHandle := adminUuid + ".test.golang"

			controllerUuid := uuid.NewV4().String()
			controllerHandle := controllerUuid + ".test.golang"

			ownerUuid := uuid.NewV4().String()
			ownerHandle := ownerUuid + ".test.golang"

			adminPrivateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			adminAddress, err := sila.GetWalletAddress(adminPrivateKey)
			So(err, ShouldBeNil)

			controllerPrivateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			controllerAddress, err := sila.GetWalletAddress(controllerPrivateKey)
			So(err, ShouldBeNil)

			ownerPrivateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			ownerAddress, err := sila.GetWalletAddress(ownerPrivateKey)
			So(err, ShouldBeNil)

			Convey("And the user handles are registered", func() {
				// In reality, the three different people with separate roles could not be the same person with the same
				// data. This is just for testing.
				homeAddress := domain.RegistrationAddress{
					AddressAlias:   "Home",
					StreetAddress1: "1234 Fake St.",
					City:           "Los Angeles",
					State:          "CA",
					Country:        "US",
					PostalCode:     "90001",
				}

				response, err := silaClient.Register(adminHandle).
					SetIndividualEntity("Alberta", "Bobbeth", "1950-10-31").
					SetAddress(homeAddress).
					SetIdentity(domain.Ssn, "181-91-1478").
					SetContact("Home", "123-456-7890", "alberta@bobbeth.com").
					SetCrypto("Main Address", adminAddress).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)

				response, err = silaClient.Register(controllerHandle).
					SetIndividualEntity("Alberta", "Bobbeth", "1950-10-31").
					SetAddress(homeAddress).
					SetIdentity(domain.Ssn, "181-91-1478").
					SetContact("Home", "123-456-7890", "alberta@bobbeth.com").
					SetCrypto("Main Address", controllerAddress).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)

				response, err = silaClient.Register(ownerHandle).
					SetIndividualEntity("Alberta", "Bobbeth", "1950-10-31").
					SetAddress(homeAddress).
					SetIdentity(domain.Ssn, "181-91-1478").
					SetContact("Home", "123-456-7890", "alberta@bobbeth.com").
					SetCrypto("Main Address", ownerAddress).
					Do()
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)

				Convey("And calls to request KYC on the newly registered accounts should succeed", func() {
					response, err := silaClient.RequestKyc(adminHandle).
						Do(adminPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)

					response, err = silaClient.RequestKyc(controllerHandle).
						Do(controllerPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)

					response, err = silaClient.RequestKyc(ownerHandle).
						Do(ownerPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)

					Convey("A call to register a business should succeed", func() {
						businessUuid := uuid.NewV4().String()
						businessHandle := businessUuid + ".business.test.golang"

						businessPrivateKey, err := sila.GenerateNewPrivateKey()
						So(err, ShouldBeNil)
						businessAddress, err := sila.GetWalletAddress(businessPrivateKey)
						So(err, ShouldBeNil)

						response, err := silaClient.Register(businessHandle).
							SetCrypto("Main Address", businessAddress).
							SetBusinessEntity("My Fancy Business", "llc", 5411).
							SetIdentity(domain.Ein, "12-1234567").
							SetAddress(homeAddress).
							SetContact("Office", "111-222-3333", "business@fancy.co").
							Do()

						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)

						Convey("A call to link the individuals to the business succeeds", func() {
							response, err := silaClient.LinkBusinessMember(adminHandle, businessHandle).
								SetAdminMember().
								Do(adminPrivateKey, businessPrivateKey)
							So(err, ShouldBeNil)
							So(response.Success, ShouldBeTrue)

							response, err = silaClient.LinkBusinessMember(adminHandle, businessHandle).
								SetControllingOfficerMemberAsAdmin(controllerHandle).
								Do(adminPrivateKey, businessPrivateKey)
							So(err, ShouldBeNil)
							So(response.Success, ShouldBeTrue)

							response, err = silaClient.LinkBusinessMember(adminHandle, businessHandle).
								SetBeneficialOwnerMemberAsAdmin(ownerHandle, 100.0).
								Do(adminPrivateKey, businessPrivateKey)
							So(err, ShouldBeNil)
							So(response.Success, ShouldBeTrue)

							Convey("A call to request the business KYC should succeed and pass.", func() {
								kycResponse, err := silaClient.RequestKyc(businessHandle).
									Do(businessPrivateKey)
								So(err, ShouldBeNil)
								So(kycResponse.Success, ShouldBeTrue)

								time.Sleep(2 * time.Minute)

								checkKycResponse, err := silaClient.CheckKyc(businessHandle).
									Do(businessPrivateKey)
								So(err, ShouldBeNil)
								So(checkKycResponse.Success, ShouldBeFalse)
								So(checkKycResponse.VerificationStatus, ShouldEqual, "passed")

								Convey("A call to get the beneficial owner entity should succeed and include the certification token.", func() {
									response, err := silaClient.GetEntity(ownerHandle).
										Do(ownerPrivateKey)
									So(err, ShouldBeNil)
									So(response.Success, ShouldBeTrue)
									So(response.UserHandle, ShouldEqual, ownerHandle)
									So(response.Memberships, ShouldHaveLength, 1)
									So(response.Memberships[0].EntityName, ShouldEqual, "My Fancy Business")
									certificationToken := response.Memberships[0].CertificationToken
									So(certificationToken, ShouldNotBeEmpty)

									Convey("A call to certify the beneficial owner entity should succeed.", func() {
										response, err := silaClient.CertifyBeneficialOwner(adminHandle, businessHandle).
											SetCertificationToken(ownerHandle, certificationToken).
											Do(adminPrivateKey, businessPrivateKey)
										So(err, ShouldBeNil)
										So(response.Success, ShouldBeTrue)

										Convey("A call to certify the business should succeed.", func() {
											response, err := silaClient.CertifyBusiness(adminHandle, businessHandle).
												Do(adminPrivateKey, businessPrivateKey)

											So(err, ShouldBeNil)
											So(response.Success, ShouldBeTrue)

											Convey("A call to check the business KYC should succeed.", func() {
												response, err := silaClient.CheckKyc(businessHandle).
													Do(businessPrivateKey)
												So(err, ShouldBeNil)
												So(response.Success, ShouldBeTrue)
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	})
}

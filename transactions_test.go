package sila_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/kappapay/silaclient"
	"github.com/kappapay/silaclient/domain"
)

func TestClient_Transactions(t *testing.T) {
	transactionSleepTime := 3*time.Minute + 30*time.Second
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
			ensureIntegrationUserExistsWithLinkedAccount(silaClient, userHandle, userWalletAddress, userWalletPrivateKey)

			Convey("A call to issue Sila coin to the main wallet should succeed", func() {
				response, err := silaClient.IssueSila(userHandle).
					SetAmountFromAccount(100, "default").
					SetDescriptor("DepositForCancel").
					SetProcessingType("STANDARD_ACH").
					SetRef("My Reference").
					Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)
				So(response.Status, ShouldEqual, "SUCCESS")
				So(response.ValidationDetails, ShouldBeNil)
				So(response.Reference, ShouldEqual, "My Reference")
				So(response.Message, ShouldEqual, "Transaction submitted to processing queue.")
				So(response.TransactionId, ShouldNotBeZeroValue)
				So(response.Descriptor, ShouldEqual, "DepositForCancel")
				transactionId := response.TransactionId

				Convey("A call to get transactions for the wallet should succeed", func() {
					now := time.Now()
					response, err := silaClient.GetTransactions(userHandle).
						SetSearchFilters(domain.TransactionSearchFilters{
							ShowTimelines: true,
							SortAscending: true,
							MaxSilaAmount: 10000000,
							MinSilaAmount: 50,
							StartTime:     now.Add(-1 * time.Hour),
							EndTime:       now.Add(1 * time.Hour),
							Page:          1,
							PerPage:       10,
						}).
						Do(userWalletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					So(response.Status, ShouldEqual, "SUCCESS")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Page, ShouldEqual, 1)
					So(response.ReturnedCount, ShouldBeGreaterThanOrEqualTo, 1)
					So(response.TotalCount, ShouldBeGreaterThanOrEqualTo, 1)
					So(response.Transactions, ShouldNotBeEmpty)
					So(response.Transactions, ShouldHaveLength, response.ReturnedCount)
				})

				Convey("A call to cancel the Sila coin issue transaction should succeed", func() {
					response, err := silaClient.CancelTransaction(userHandle, transactionId).
						SetRef("My Reference").
						Do(userWalletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					So(response.Status, ShouldEqual, "SUCCESS")
					So(response.ValidationDetails, ShouldBeNil)
					So(response.Reference, ShouldEqual, "My Reference")

					Convey("The transaction should display as cancelled when fetched", func() {
						response, err := silaClient.GetTransactions(userHandle).
							SetSearchFilters(domain.TransactionSearchFilters{
								TransactionId: transactionId,
								Statuses:      []string{"failed"},
								Page:          1,
								PerPage:       5,
							}).
							Do(userWalletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.ReturnedCount, ShouldNotBeEmpty)
					})
				})
			})

			Convey("And the specified integration user has a second wallet", func() {
				newWalletPrivateKey, err := sila.GenerateNewPrivateKey()
				So(err, ShouldBeNil)
				newWalletAddress, err := sila.GetWalletAddress(newWalletPrivateKey)
				So(err, ShouldBeNil)
				signature, err := sila.GenerateWalletSignature([]byte(newWalletAddress), newWalletPrivateKey)

				response, err := silaClient.RegisterWallet(userHandle).
					SetWallet("transaction test wallet", newWalletAddress, signature).
					Do(userWalletPrivateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldBeTrue)

				Convey("A call to issue Sila coin to the main wallet should succeed", func() {
					response, err := silaClient.IssueSila(userHandle).
						SetAmountFromAccount(100, "default").
						SetDescriptor("DepositToSilaWallet").
						SetProcessingType("STANDARD_ACH").
						Do(userWalletPrivateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldBeTrue)
					transactionId := response.TransactionId

					Convey("Wait for a bit, then check to see that the issue succeeded", func() {
						time.Sleep(transactionSleepTime)
						response, err := silaClient.GetTransactions(userHandle).
							SetSearchFilters(domain.TransactionSearchFilters{
								TransactionId: transactionId,
								Page:          1,
								PerPage:       1,
							}).
							Do(userWalletPrivateKey)
						So(err, ShouldBeNil)
						So(response.Success, ShouldBeTrue)
						So(response.ReturnedCount, ShouldNotBeEmpty)
						So(response.Transactions[0].Status, ShouldEqual, "success")

						Convey("A call to transfer Sila coin from the main wallet to the secondary wallet should succeed", func() {
							response, err := silaClient.TransferSila(userHandle).
								SetAmountAndUser(100, userHandle).
								SetDestinationAddress(newWalletAddress).
								SetDescriptor("Moving Money").
								SetRef("My Reference").
								Do(userWalletPrivateKey)
							So(err, ShouldBeNil)
							So(response.Success, ShouldBeTrue)
							So(response.Status, ShouldEqual, "SUCCESS")
							So(response.ValidationDetails, ShouldBeNil)
							So(response.Reference, ShouldEqual, "My Reference")
							So(response.Message, ShouldEqual, "Transaction submitted to processing queue.")
							So(response.TransactionId, ShouldNotBeZeroValue)
							So(response.Descriptor, ShouldEqual, "Moving Money")
							transactionId := response.TransactionId

							Convey("Wait for a bit, then check to see that the transfer succeeded", func() {
								time.Sleep(transactionSleepTime)
								response, err := silaClient.GetTransactions(userHandle).
									SetSearchFilters(domain.TransactionSearchFilters{
										TransactionId: transactionId,
										Page:          1,
										PerPage:       1,
									}).
									Do(userWalletPrivateKey)
								So(err, ShouldBeNil)
								So(response.Success, ShouldBeTrue)
								So(response.ReturnedCount, ShouldNotBeEmpty)
								So(response.Transactions[0].Status, ShouldEqual, "success")

								Convey("A call to redeem Sila coin from the secondary wallet should succeed", func() {
									response, err := silaClient.RedeemSila(userHandle).
										SetAmountToAccount(100, "default").
										SetDescriptor("Redeem Money").
										SetProcessingType("STANDARD_ACH").
										SetRef("My Reference").
										Do(newWalletPrivateKey)
									So(err, ShouldBeNil)
									So(response.Success, ShouldBeTrue)
									So(response.Status, ShouldEqual, "SUCCESS")
									So(response.ValidationDetails, ShouldBeNil)
									So(response.Reference, ShouldEqual, "My Reference")
									So(response.Message, ShouldEqual, "Transaction submitted to processing queue.")
									So(response.TransactionId, ShouldNotBeZeroValue)
									So(response.Descriptor, ShouldEqual, "Redeem Money")
									transactionId := response.TransactionId

									Convey("Wait for a bit, then check to see that the transfer succeeded", func() {
										time.Sleep(transactionSleepTime)
										response, err := silaClient.GetTransactions(userHandle).
											SetSearchFilters(domain.TransactionSearchFilters{
												TransactionId: transactionId,
												Page:          1,
												PerPage:       1,
											}).
											Do(userWalletPrivateKey)
										So(err, ShouldBeNil)
										So(response.Success, ShouldBeTrue)
										So(response.ReturnedCount, ShouldNotBeEmpty)
										So(response.Transactions[0].Status, ShouldEqual, "success")
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

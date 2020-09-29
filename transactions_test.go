package sila_test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_IssueSila(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to issue Sila coin to a wallet should succeed", func() {
				response, err := client.IssueSila("user.silamoney.eth").
					SetAmountToAccount(10000, "default").
					SetDescriptor("RentUnit#7").
					SetProcessingType("STANDARD_ACH").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_TransferSila(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to transfer Sila coin from one wallet to another should succeed", func() {
				response, err := client.TransferSila("user.silamoney.eth").
					SetAmountAndUser(10000, "user2.silamoney.eth").
					SetDescriptor("RentUnit#7").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_RedeemSila(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to redeem Sila coin from a wallet should succeed", func() {
				response, err := client.RedeemSila("user.silamoney.eth").
					SetAmountFromAccount(10000, "default").
					SetDescriptor("RentUnit#7").
					SetProcessingType("STANDARD_ACH").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_GetTransactions(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("The call to get transactions should succeed", func() {
				now := time.Now()
				response, err := client.GetTransactions("user.silamoney.eth").
					SetSearchFilters(sila.TransactionSearchFilters{
						ShowTimelines: true,
						SortAscending: true,
						MaxSilaAmount: 10000000,
						MinSilaAmount: 50,
						StartTime:     now.Add(-4 * time.Hour),
						EndTime:       now.Add(4 * time.Hour),
						Page:          1,
						PerPage:       10,
					}).
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(response.Success, ShouldEqual, "SUCCESS")
			})
		})
	})
}

func TestClient_CancelTransaction(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)
			Convey("A transaction to issue sila coin should exist and be in progress", func() {
				transactionResponse, err := client.IssueSila("user.silamoney.eth").
					SetAmountToAccount(10000, "default").
					SetDescriptor("RentUnit#7").
					SetProcessingType("STANDARD_ACH").
					SetRef("My Reference").
					Do(privateKey)
				So(err, ShouldBeNil)
				So(transactionResponse.Success, ShouldEqual, "SUCCESS")

				Convey("The call to cancel a transaction should succeed", func() {
					response, err := client.CancelTransaction("user.silamoney.eth", transactionResponse.TransactionId).
						SetRef("My Reference").
						Do(privateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldEqual, "SUCCESS")
				})
			})
		})
	})
}

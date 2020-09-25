package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_RegisterWallet(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The 'existing' wallet private key and address are generated", func() {
			privateKey, err := sila.GenerateNewPrivateKey()
			So(err, ShouldBeNil)

			Convey("The new wallet private key and address are generated, along with a signature", func() {
				newWalletPrivateKey, err := sila.GenerateNewPrivateKey()
				So(err, ShouldBeNil)
				address, err := sila.GetWalletAddress(privateKey)
				So(err, ShouldBeNil)
				signature, err := sila.GenerateWalletSignature([]byte(address), newWalletPrivateKey)

				Convey("The call to register a wallet should succeed", func() {
					response, err := client.RegisterWallet("user.silamoney.eth").
						SetRef("My Reference").
						SetWallet("My Wallet", address, signature).
						Do(privateKey)
					So(err, ShouldBeNil)
					So(response.Success, ShouldEqual, "SUCCESS")
				})
			})
		})
	})
}

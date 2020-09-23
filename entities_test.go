package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_CheckHandle(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to check handle should succeed", func() {
			response, err := client.CheckHandle("user.silamoney.eth").Ref("My Reference").Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldEqual, "SUCCESS")
		})
	})
}

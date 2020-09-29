package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_GetBusinessTypes(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to get business types should succeed", func() {
			response, err := client.GetBusinessTypes().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldEqual, "SUCCESS")
		})

	})
}

func TestClient_GetBusinessRoles(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to get business roles should succeed", func() {
			response, err := client.GetBusinessRoles().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldEqual, "SUCCESS")
		})

	})
}

func TestClient_GetNaicsCategories(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		client, err := sila.NewClient(
			"badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c",
			"handle.silamoney.eth",
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to get NAICS categories should succeed", func() {
			response, err := client.GetNaicsCategories().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldEqual, "SUCCESS")
		})

	})
}

package sila_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"sila"
)

func TestClient_Parameters(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := ReadTestConfig()
		So(err, ShouldBeNil)
		client, err := sila.NewClient(
			testConfig.AuthPrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to get business types should succeed", func() {
			response, err := client.GetBusinessTypes().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
			So(response.BusinessTypes, ShouldNotBeEmpty)
		})
		Convey("The call to get business roles should succeed", func() {
			response, err := client.GetBusinessRoles().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
			So(response.BusinessRoles, ShouldNotBeEmpty)
		})
		Convey("The call to get NAICS categories should succeed", func() {
			response, err := client.GetNaicsCategories().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
			So(response.NaicsCategories, ShouldNotBeEmpty)
		})
	})
}

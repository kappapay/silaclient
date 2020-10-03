package sila_test

import (
	"sila"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClient_Parameters(t *testing.T) {
	Convey("Given the Sila client exists", t, func() {
		testConfig, err := readTestConfig()
		So(err, ShouldBeNil)
		silaClient, err := sila.NewClient(
			testConfig.AuthPrivateKeyKex,
			testConfig.AuthHandle,
			sila.Sandbox)
		So(err, ShouldBeNil)
		Convey("The call to get business types should succeed", func() {
			response, err := silaClient.GetBusinessTypes().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
			So(response.BusinessTypes, ShouldNotBeEmpty)
		})
		Convey("The call to get business roles should succeed", func() {
			response, err := silaClient.GetBusinessRoles().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
			So(response.BusinessRoles, ShouldNotBeEmpty)
		})
		Convey("The call to get NAICS categories should succeed", func() {
			response, err := silaClient.GetNaicsCategories().Do()
			So(err, ShouldBeNil)
			So(response.Success, ShouldBeTrue)
			So(response.Status, ShouldEqual, "SUCCESS")
			So(response.ValidationDetails, ShouldBeNil)
			So(response.NaicsCategories, ShouldNotBeEmpty)
		})
	})
}

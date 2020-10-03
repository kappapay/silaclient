package sila_test

import (
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"

	"sila"
	"sila/domain"
)

type TestConfigData struct {
	AuthPrivateKeyKex       string
	AuthHandle              string
	UserHandle              string
	UserWalletPrivateKeyHex string
}

func readTestConfig() (TestConfigData, error) {
	testConfigData := TestConfigData{}
	config := viper.New()
	config.SetConfigName("test_config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		return testConfigData, err
	}
	testConfigData.AuthPrivateKeyKex = config.GetString("auth.private_key_hex")
	testConfigData.AuthHandle = config.GetString("auth.handle")
	testConfigData.UserHandle = config.GetString("user.handle")
	testConfigData.UserWalletPrivateKeyHex = config.GetString("user.wallet_private_key_hex")
	return testConfigData, nil
}

func ensureIntegrationUserExistsWithLinkedAccount(silaClient sila.Client, userHandle string, userWalletAddress string, userWalletPrivateKey string) {
	ensureIntegrationUserExists(silaClient, userHandle, userWalletAddress, userWalletPrivateKey)
	response, err := silaClient.LinkAccount(userHandle).
		SetAccountName("default").
		SetAccountType("CHECKING").
		SetDirectLinkAccount("123456789012", "123456789").
		Do(userWalletPrivateKey)
	So(err, ShouldBeNil)
	So(response.Success, ShouldBeTrue)
}

func ensureIntegrationUserExists(silaClient sila.Client, userHandle string, userWalletAddress string, userWalletPrivateKey string) {
	response, err := silaClient.CheckHandle(userHandle).Do()
	So(err, ShouldBeNil)
	if response.Success == true {
		registerResponse, err := silaClient.Register(userHandle).
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
			SetCrypto("Main Address", userWalletAddress).
			Do()
		So(err, ShouldBeNil)
		So(registerResponse.Success, ShouldBeTrue)

		requestKycResponse, err := silaClient.RequestKyc(userHandle).Do(userWalletPrivateKey)
		So(err, ShouldBeNil)
		So(requestKycResponse.Success, ShouldBeTrue)

		time.Sleep(30 * time.Second)

		checkKycResponse, err := silaClient.CheckKyc(userHandle).Do(userWalletPrivateKey)
		So(err, ShouldBeNil)
		So(checkKycResponse.Success, ShouldBeTrue)
	}
}

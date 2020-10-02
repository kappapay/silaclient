package sila_test

import (
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"

	"sila"
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

func ensureIntegrationUserExistsWithLinkedAccount(client *sila.Client, userHandle string, userWalletAddress string, userWalletPrivateKey string) {
	ensureIntegrationUserExists(client, userHandle, userWalletAddress, userWalletPrivateKey)
	response, err := client.LinkAccount(userHandle).
		SetAccountName("default").
		SetAccountType("CHECKING").
		SetDirectLinkAccount("123456789012", "123456789").
		Do(userWalletPrivateKey)
	So(err, ShouldBeNil)
	So(response.Success, ShouldBeTrue)
}

func ensureIntegrationUserExists(client *sila.Client, userHandle string, userWalletAddress string, userWalletPrivateKey string) {
	response, err := client.CheckHandle(userHandle).Do()
	So(err, ShouldBeNil)
	if response.Success == true {
		registerResponse, err := client.Register(userHandle).
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
			SetCrypto("Main Address", userWalletAddress).
			Do()
		So(err, ShouldBeNil)
		So(registerResponse.Success, ShouldBeTrue)

		requestKycResponse, err := client.RequestKyc(userHandle).Do(userWalletPrivateKey)
		So(err, ShouldBeNil)
		So(requestKycResponse.Success, ShouldBeTrue)

		time.Sleep(30 * time.Second)

		checkKycResponse, err := client.CheckKyc(userHandle).Do(userWalletPrivateKey)
		So(err, ShouldBeNil)
		So(checkKycResponse.Success, ShouldBeTrue)
	}
}

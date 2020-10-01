package sila_test

import "github.com/spf13/viper"

type TestConfigData struct {
	AuthPrivateKeyKex       string
	AuthHandle              string
	UserHandle              string
	UserWalletPrivateKeyHex string
}

func ReadTestConfig() (TestConfigData, error) {
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

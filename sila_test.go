package sila_test

import "github.com/spf13/viper"

type TestConfigData struct {
	PrivateKeyKex string
	AuthHandle    string
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
	testConfigData.PrivateKeyKex = config.GetString("private_key_hex")
	testConfigData.AuthHandle = config.GetString("auth_handle")
	return testConfigData, nil
}

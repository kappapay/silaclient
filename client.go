package sila

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Client struct {
	privateKey  *ecdsa.PrivateKey
	authHandle  string
	version     string
	crypto      string
	environment Environment
}

type Environment string

const (
	Sandbox    Environment = "https://sandbox.silamoney.com/"
	Production             = "https://api.silamoney.com/"
)

func (env Environment) generateURL(version string, path string) string {
	return string(env) + version + path
}

func NewClient(privateKeyHex string, authHandle string, environment Environment) (*Client, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.Errorf("private key invalid, make sure it is hex without the 0x prefix: %v", err)
	}
	return &Client{
		privateKey:  privateKey,
		authHandle:  authHandle,
		version:     "0.2",
		crypto:      "ETH",
		environment: environment,
	}, nil
}

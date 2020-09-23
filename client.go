package sila

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"sync"

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

var once sync.Once

var (
	instance *Client
)

func (env Environment) generateURL(version string, path string) string {
	return string(env) + version + path
}

func NewClient(privateKeyHex string, authHandle string, environment Environment) (*Client, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.Errorf("private key invalid, make sure it is hex without the 0x prefix: %v", err)
	}
	once.Do(func() {
		instance = &Client{
			privateKey:  privateKey,
			authHandle:  authHandle,
			version:     "0.2",
			crypto:      "ETH",
			environment: environment,
		}
	})
	return instance, nil
}

func (client Client) GenerateAuthSignature(requestBody []byte) string {
	// Follows the Sila example for Golang

	// Generate the message hash using the Keccak 256 algorithm.
	msgHash := crypto.Keccak256(requestBody)

	// Create a signature using your private key and hashed message.
	sigBytes, err := crypto.Sign(msgHash, client.privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// The signature just created is off by -27 from what the API
	// will expect. Correct that by converting the signature bytes
	// to a big int and adding 27.
	var offset int64 = 27
	var bigSig = new(big.Int).SetBytes(sigBytes)
	sigBytes = bigSig.Add(bigSig, big.NewInt(offset)).Bytes()

	// The big library takes out any padding, but the resultant
	// signature must be 130 characters (65 bytes) long. In some
	// cases, you might find that sigBytes now has a length of 64 or
	// less, so you can fix that in this way (this prepends the hex
	// value with "0" until the requisite length is reached).
	// Example: if two digits were required but the value was 1, you'd
	// pass in 01.
	var sigBytesLength = 65 // length of a valid signature byte array
	var arr = make([]byte, sigBytesLength)
	copy(arr[(sigBytesLength-len(sigBytes)):], sigBytes)

	// Encode the bytes to a hex string.
	return hex.EncodeToString(arr)
}

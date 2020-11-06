package sila

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

type Client interface {
	// Determine if a handle is currently on use in the network
	CheckHandle(userHandle string) CheckHandle
	// Register an individual or business user to a handle, add basic KYC data, and an initial wallet
	Register(userOrBusinessHandle string) Register
	// Request KYC for an individual user
	RequestKyc(userHandle string) RequestKyc
	// Check the status of KYC for an individual user
	CheckKyc(userHandle string) CheckKyc
	// Gets an information about a user
	GetEntity(userHandle string) GetEntity
	// Searches for user information
	GetEntities() GetEntities
	// Links an individual user to a business user, indicating the individual is part of that business
	LinkBusinessMember(userHandle string, businessHandle string) LinkBusinessMember
	// Unlinks an individual user from a business user, indicating that they are not part of that business
	UnlinkBusinessMember(userHandle string, businessHandle string) UnlinkBusinessMember
	// Have a business's admin user certify that the information about a beneficial owner (who has ownership stake in the business) has correct information
	CertifyBeneficialOwner(adminUserHandle string, businessHandle string) CertifyBeneficialOwner
	// Have a business's admin user certify that all information about a business has been entered correctly and the business should be allowed to transact on the network
	CertifyBusiness(adminUserHandle string, businessHandle string) CertifyBusiness
	// Add registration data to a user after registration
	AddRegistrationData(userHandle string) AddRegistrationData
	// Update a user's registration data after registration
	UpdateRegistrationData(userHandle string) UpdateRegistrationData
	// Delete a user's registration data after registration
	DeleteRegistrationData(userHandle string) DeleteRegistrationData

	// Link a bank account to a user, either directly or via Plaid
	LinkAccount(userHandle string) LinkAccount
	// Complete same day auth with Plaid
	PlaidSameDayAuth(userHandle string, accountName string) PlaidSameDayAuth
	// Get a user's linked accounts
	GetAccounts(userHandle string) GetAccounts
	// Get the balances for a user's linked accounts
	GetAccountBalance(userHandle string, accountName string) GetAccountBalance

	// Register a new wallet to a user
	RegisterWallet(userHandle string) RegisterWallet
	// Get a user's wallet
	GetWallet(userHandle string) GetWallet
	// Get several of a user's wallets
	GetWallets(userHandle string) GetWallets
	// Update information about a user's wallet
	UpdateWallet(userHandle string) UpdateWallet
	// Get the current Sila coin balance of a user's wallet
	GetWalletBalance(walletAddress string) GetSilaBalance
	// Delete a user's wallet
	DeleteWallet(userHandle string) DeleteWallet

	// Issue Sila coin to a wallet from a linked bank account
	IssueSila(userHandle string) IssueSila
	// Transfer Sila coin from one wallet to another
	TransferSila(userHandle string) TransferSila
	// Redeem Sila coin from a user's wallet to a linked bank account
	RedeemSila(userHandle string) RedeemSila
	// Get a list of transactions related to a user's wallet
	GetTransactions(userHandle string) GetTransactions
	// Cancel a user's transaction if still in progress
	CancelTransaction(userHandle string, transactionId string) CancelTransactions

	// Get a list of business types
	GetBusinessTypes() GetBusinessTypes
	// Get a list of business roles
	GetBusinessRoles() GetBusinessRoles
	// Get a list of NAICS categories and their codes
	GetNaicsCategories() GetNaicsCategories
}

// The Sila client for handling calls to the Sila API
type ClientImpl struct {
	privateKey  *ecdsa.PrivateKey
	authHandle  string
	version     string
	crypto      string
	environment Environment
}

// Which API environment to run in
type Environment string

const (
	Sandbox    Environment = "https://sandbox.silamoney.com/"
	Production             = "https://api.silamoney.com/"
)

var once sync.Once

var (
	// A singleton instance for the client
	instance *ClientImpl
)

// Generates a URL for the current environment given the API version and the path to invoke
func (env Environment) generateURL(version string, path string) string {
	return string(env) + version + path
}

// Creates a new Sila client using your system's auth private key as a hex string, your system's auth handle, and the
// environment to send requests to (sandbox or production).
func NewClient(privateKeyHex string, authHandle string, environment Environment) (Client, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, errors.Errorf("private key invalid, make sure it is hex without the 0x prefix: %v", err)
	}
	once.Do(func() {
		instance = &ClientImpl{
			privateKey:  privateKey,
			authHandle:  authHandle,
			version:     "0.2",
			crypto:      "ETH",
			environment: environment,
		}
	})
	return instance, nil
}

// Generates a signature for a request's body using the provided private key.
func generateSignatureFromKey(requestBody []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	// Follows the Sila example for Golang
	// Generate the message hash using the Keccak 256 algorithm.
	msgHash := crypto.Keccak256(requestBody)

	// Create a signature using your private key and hashed message.
	sigBytes, err := crypto.Sign(msgHash, privateKey)
	if err != nil {
		return "", err
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
	return hex.EncodeToString(arr), nil
}

// Generates a signature for a message with your system's private auth key from the client creation.
func (client ClientImpl) generateAuthSignature(message []byte) (string, error) {
	return generateSignatureFromKey(message, client.privateKey)
}

// Perform a call to the API at some path with the included request and a pointer to the response struct
func (client *ClientImpl) performCall(path string, requestBody interface{}, responseBody interface{}) error {
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil
	}
	url := instance.environment.generateURL(instance.version, path)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return err
	}
	request.Header.Set("Content-type", "application/json")
	authSignature, err := instance.generateAuthSignature(requestJson)
	if err != nil {
		return errors.Errorf("failed to generate auth signature: %v", err)
	}
	request.Header.Set("authsignature", authSignature)
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	return nil
}

// Perform a call to the API at some path signed by a user's wallet private key, with the included request and a pointer to the response struct
func (client *ClientImpl) performCallWithUserAuth(path string, requestBody interface{}, responseBody interface{}, userWalletPrivateKey string) error {
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil
	}
	url := instance.environment.generateURL(instance.version, path)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return err
	}
	request.Header.Set("Content-type", "application/json")
	authSignature, err := instance.generateAuthSignature(requestJson)
	if err != nil {
		return errors.Errorf("failed to generate auth signature: %v", err)
	}
	request.Header.Set("authsignature", authSignature)
	userSignature, err := GenerateWalletSignature(requestJson, userWalletPrivateKey)
	if err != nil {
		return errors.Errorf("failed to generate user signature: %v", err)
	}
	request.Header.Set("usersignature", userSignature)
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	return nil
}

// Perform a call to the API at some path signed by a user's wallet private key and a business's wallet key, with the included request and a pointer to the response struct
func (client *ClientImpl) performCallWithUserAndBusinessAuth(path string, requestBody interface{}, responseBody interface{}, userWalletPrivateKey string, businessWalletPrivateKey string) error {
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil
	}
	url := instance.environment.generateURL(instance.version, path)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return err
	}
	request.Header.Set("Content-type", "application/json")

	authSignature, err := instance.generateAuthSignature(requestJson)
	if err != nil {
		return errors.Errorf("failed to generate auth signature: %v", err)
	}
	request.Header.Set("authsignature", authSignature)

	userSignature, err := GenerateWalletSignature(requestJson, userWalletPrivateKey)
	if err != nil {
		return errors.Errorf("failed to generate user signature: %v", err)
	}
	request.Header.Set("usersignature", userSignature)

	businessSignature, err := GenerateWalletSignature(requestJson, businessWalletPrivateKey)
	if err != nil {
		return errors.Errorf("failed to generate business signature: %v", err)
	}
	request.Header.Set("businesssignature", businessSignature)

	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	return nil
}

// Perform a public (no auth required) call to the API at some path with the included request and a pointer to the response struct
func (client *ClientImpl) performPublicCall(path string, requestBody interface{}, responseBody interface{}) error {
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return nil
	}
	url := instance.environment.generateURL(instance.version, path)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return err
	}
	request.Header.Set("Content-type", "application/json")
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return err
	}
	return nil
}

// Gets a wallet address from a wallet's private key as a hex string and returns the wallet address
func GetWalletAddress(privateKeyHex string) (string, error) {
	publicKeyECDSA, err := getPublicKeyFromPrivateHex(privateKeyHex)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
}

// Gets a public key from a private key hex string
func getPublicKeyFromPrivateHex(privateKeyHex string) (*ecdsa.PublicKey, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	return publicKeyECDSA, nil
}

// Generates a new private key as a hex string for a wallet
func GenerateNewPrivateKey() (string, error) {
	pk, err := crypto.GenerateKey()
	if err != nil {
		return "", err
	}
	pkBytes := crypto.FromECDSA(pk)
	pkHex := hexutil.Encode(pkBytes)[2:]
	return pkHex, nil
}

// Generates a signature for a message with one of a user's wallet private keys (in hex) as provided.
func GenerateWalletSignature(message []byte, walletPrivateKeyHex string) (string, error) {
	privateKey, err := crypto.HexToECDSA(walletPrivateKeyHex)
	if err != nil {
		return "", errors.Errorf("private key invalid, make sure it is hex without the 0x prefix: %v", err)
	}
	return generateSignatureFromKey(message, privateKey)
}

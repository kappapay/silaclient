package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/ardanlabs/conf/v3"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	sila "github.com/kappapay/silaclient"
	"github.com/kappapay/silaclient/domain"
)

type SilaConfig struct {
	PrivateKeyHex string `conf:"env:SILA_PRIVATE_KEY_HEX,required"`
	AuthHandle    string `conf:"env:SILA_AUTH_HANDLE,required"`
	BaseURL       string `conf:"env:SILA_BASE_URL,required"`
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			return fmt.Errorf("error loading .env file: %w", err)
		}
	}

	cfg := SilaConfig{}
	_, err := conf.Parse("", &cfg)
	if err != nil {
		return fmt.Errorf("parsing config: %w", err)
	}
	client, err := sila.NewClient(cfg.PrivateKeyHex, cfg.AuthHandle, sila.Environment(cfg.BaseURL))
	if err != nil {
		return fmt.Errorf("error initializing client: %w", err)
	}

	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("failed to get public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	addressHex := address.Hex()

	userHandle := uuid.NewString()
	resp, err := client.Register(userHandle).
		SetAddress(domain.RegistrationAddress{
			StreetAddress1: "123 Brick St.",
			City:           "Hometown",
			PostalCode:     "12345",
			State:          "WY",
		}).
		SetContact("default", "+13335554444", "example@example.com").
		SetCrypto("default", addressHex).
		SetIndividualEntity("Don", "Wurribouttet", "1970-04-20").
		SetIdentity(domain.Ssn, "333224444").
		Do()
	if err != nil {
		return err
	}

	if !resp.Success {
		return fmt.Errorf("sila register call failed: %v", resp)
	}

	userPrivateKeyHex := hex.EncodeToString(privateKey.D.Bytes())
	kycResp, err := client.RequestKyc(userHandle).Do(userPrivateKeyHex)
	if err != nil {
		return err
	}

	fmt.Println("request_kyc:", kycResp)

	kycResultResp, err := client.CheckKyc(userHandle).Do(userPrivateKeyHex)
	if err != nil {
		return err
	}
	fmt.Println("kyc_result:", kycResultResp)

	fmt.Println("user address:", addressHex)
	fmt.Printf("USER_HANDLE=%s\n", userHandle)
	fmt.Printf("USER_PRIVATE_KEY_HEX=%s\n", userPrivateKeyHex)
	return nil

}

package sila

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type Header struct {
	Created    int64  `json:"created"`
	AuthHandle string `json:"auth_handle"`
	UserHandle string `json:"user_handle"`
	Version    string `json:"version"`
	Crypto     string `json:"crypto"`
	Reference  string `json:"reference,omitempty"`
}

func (client Client) generateHeader(userHandle string, reference string) Header {
	return Header{
		Created:    time.Now().Unix(),
		AuthHandle: client.authHandle,
		UserHandle: userHandle,
		Version:    client.version,
		Crypto:     client.crypto,
		Reference:  reference,
	}
}

func (header *Header) Ref(ref string) {
	header.Reference = ref
}

type CheckHandleRequest struct {
	Header  Header `json:"header"`
	Message string `json:"message"`
}

type CheckHandleResponse struct {
	Success   bool   `json:"success"`
	Reference string `json:"reference"`
	Message   string `json:"message"`
	Status    string `json:"status"`
}

func (client Client) CheckHandle(userHandle string) (CheckHandleResponse, error) {
	var responseBody CheckHandleResponse
	url := client.environment.generateURL(client.version, "/check_handle")
	requestBody := CheckHandleRequest{
		Header:  client.generateHeader(userHandle, ""),
		Message: "header_msg",
	}
	requestJson, err := json.Marshal(requestBody)
	if err != nil {
		return responseBody, nil
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(requestJson))
	if err != nil {
		return responseBody, err
	}
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("authsignature", client.GenerateAuthSignature(requestJson))
	httpClient := http.Client{}
	resp, err := httpClient.Do(request)
	if err != nil {
		return responseBody, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return responseBody, err
	}
	return responseBody, nil
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

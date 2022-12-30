package sila

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"

	"github.com/kappapay/silaclient/domain"
)

type DocumentUpload interface {
	AddFile(metadata FileMetadata, fileContents []byte) DocumentUpload
	Do(userWalletPrivateKey string) (domain.DocumentUploadResponse, error)
}

type DocumentUploadMsg struct {
	Files map[string][]byte  `json:"-"`
	Data  DocumentUploadData `json:"data"`
}

var _ DocumentUpload = &DocumentUploadMsg{}

func (client *ClientImpl) Documents(userHandle string) DocumentUpload {
	return &DocumentUploadMsg{
		Files: map[string][]byte{},
		Data: DocumentUploadData{

			Header:       client.generateHeader().setUserHandle(userHandle),
			FileMetadata: map[string]FileMetadata{},
		},
	}
}

type FileMetadata struct {
	Name         string `json:"name"`
	Filename     string `json:"filename"`
	Hash         string `json:"hash"`
	MimeType     string `json:"mime_type"`
	DocumentType string `json:"document_type"`
	Description  string `json:"description"`
}

type DocumentUploadData struct {
	Header       *Header                 `json:"header"`
	FileMetadata map[string]FileMetadata `json:"file_metadata"`
}

// AddFile ...
func (msg *DocumentUploadMsg) AddFile(metadata FileMetadata, fileContents []byte) DocumentUpload {
	fileID := uuid.NewString()

	h := sha256.New()
	h.Write([]byte(fileContents))
	fileHash := hex.EncodeToString(h.Sum(nil))
	metadata.Hash = string(fileHash)

	msg.Files[fileID] = fileContents
	msg.Data.FileMetadata[fileID] = metadata
	return msg
}

func (msg *DocumentUploadMsg) Do(userWalletPrivateKey string) (domain.DocumentUploadResponse, error) {
	var responseBody domain.DocumentUploadResponse
	err := instance.performCallMultipartWithUserAuth("/documents", msg.Files, msg.Data, &responseBody, userWalletPrivateKey)
	return responseBody, err
}

package sila

import "time"

type Header struct {
	Created    int64  `json:"created"`
	AuthHandle string `json:"auth_handle"`
	UserHandle string `json:"user_handle,omitempty"`
	Version    string `json:"version"`
	Crypto     string `json:"crypto"`
	Reference  string `json:"reference,omitempty"`
}

func (client Client) generateHeader() *Header {
	return &Header{
		Created:    time.Now().Unix(),
		AuthHandle: client.authHandle,
		Version:    client.version,
		Crypto:     client.crypto,
	}
}

func (header *Header) setUserHandle(userHandle string) *Header {
	header.UserHandle = userHandle
	return header
}

func (header *Header) setRef(ref string) *Header {
	header.Reference = ref
	return header
}

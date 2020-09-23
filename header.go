package sila

import "time"

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

func (header *Header) Ref(ref string) *Header {
	header.Reference = ref
	return header
}

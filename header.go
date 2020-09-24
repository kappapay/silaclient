package sila

import "time"

// The header segment for every request made to Sila
type Header struct {
	Created    int64  `json:"created"`
	AuthHandle string `json:"auth_handle"`
	UserHandle string `json:"user_handle,omitempty"`
	Version    string `json:"version"`
	Crypto     string `json:"crypto"`
	Reference  string `json:"reference,omitempty"`
}

// Generate a header given information that can be readily generated or inferred from the Sila client
func (client Client) generateHeader() *Header {
	return &Header{
		Created:    time.Now().Unix(),
		AuthHandle: client.authHandle,
		Version:    client.version,
		Crypto:     client.crypto,
	}
}

// Sets the user handle in the header
func (header *Header) setUserHandle(userHandle string) *Header {
	header.UserHandle = userHandle
	return header
}

// Sets a reference message in the header which will be repeated back on the response
func (header *Header) setRef(ref string) *Header {
	header.Reference = ref
	return header
}

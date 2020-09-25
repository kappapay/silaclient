package sila

func (client Client) RegisterWallet(userHandle string) *RegisterWallet {
	return &RegisterWallet{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

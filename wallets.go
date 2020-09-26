package sila

func (client Client) RegisterWallet(userHandle string) *RegisterWallet {
	return &RegisterWallet{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

func (client Client) GetWallet(userHandle string) *GetWallet {
	return &GetWallet{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

func (client Client) GetWallets(userHandle string) *GetWallets {
	return &GetWallets{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

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

func (client Client) UpdateWallet(userHandle string) *UpdateWallet {
	return &UpdateWallet{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

func (client Client) GetWalletBalance(walletAddress string) *GetSilaBalance {
	return &GetSilaBalance{
		Address: walletAddress,
	}
}

func (client Client) DeleteWallet(userHandle string) *DeleteWallet {
	return &DeleteWallet{
		Header: client.generateHeader().setUserHandle(userHandle),
	}
}

package sila

func (client Client) LinkAccount(userHandle string) *LinkAccount {
	return &LinkAccount{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "link_account_msg",
	}
}

func (client Client) PlaidSameDayAuth(userHandle string, accountName string) *PlaidSameDayAuth {
	return &PlaidSameDayAuth{
		Header:      client.generateHeader().setUserHandle(userHandle),
		AccountName: accountName,
	}
}

func (client Client) GetAccounts(userHandle string) *GetAccounts {
	return &GetAccounts{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "get_accounts_msg",
	}
}

func (client Client) GetAccountBalance(userHandle string, accountName string) *GetAccountBalance {
	return &GetAccountBalance{
		Header:      client.generateHeader().setUserHandle(userHandle),
		AccountName: accountName,
	}
}

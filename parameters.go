package sila

func (client Client) GetBusinessTypes() *GetBusinessTypes {
	return &GetBusinessTypes{
		Header: client.generateHeader(),
	}
}

func (client Client) GetBusinessRoles() *GetBusinessRoles {
	return &GetBusinessRoles{
		Header: client.generateHeader(),
	}
}

func (client Client) GetNaicsCategories() *GetNaicsCategories {
	return &GetNaicsCategories{
		Header: client.generateHeader(),
	}
}

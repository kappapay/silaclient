package sila

func (client Client) CheckHandle(userHandle string) *CheckHandle {
	return &CheckHandle{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

func (client Client) Register(userHandle string) *Register {
	return &Register{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "entity_msg",
	}
}

func (client Client) RequestKyc(userHandle string) *RequestKyc {
	return &RequestKyc{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

func (client Client) CheckKyc(userHandle string) *CheckKyc {
	return &CheckKyc{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

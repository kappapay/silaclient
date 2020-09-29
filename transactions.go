package sila

func (client Client) IssueSila(userHandle string) *IssueSila {
	return &IssueSila{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "issue_msg",
	}
}

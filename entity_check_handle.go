package sila

import (
	"github.com/bpancost/sila/domain"
)

func (client ClientImpl) CheckHandle(userHandle string) CheckHandle {
	return &CheckHandleMsg{
		Header:  client.generateHeader().setUserHandle(userHandle),
		Message: "header_msg",
	}
}

type CheckHandleMsg struct {
	Header  *Header `json:"header"`
	Message string  `json:"message"`
}

func (msg *CheckHandleMsg) SetRef(ref string) CheckHandle {
	msg.Header.setRef(ref)
	return msg
}

func (msg *CheckHandleMsg) Do() (domain.SuccessResponse, error) {
	var responseBody domain.SuccessResponse
	err := instance.performCall("/check_handle", msg, &responseBody)
	return responseBody, err
}

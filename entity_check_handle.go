package sila

type CheckHandle struct {
	Header  *Header `json:"header"`
	Message string  `json:"message"`
}

func (msg *CheckHandle) SetRef(ref string) *CheckHandle {
	msg.Header.setRef(ref)
	return msg
}

func (msg *CheckHandle) Do() (SuccessResponse, error) {
	var responseBody SuccessResponse
	err := instance.performCall("/check_handle", msg, &responseBody)
	return responseBody, err
}

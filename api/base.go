package api

import "encoding/json"

type Request interface {
	Params() map[string]string
}

type Response interface {
	Unmarshal([]byte) error
}

type Api interface {
	Name() string
	GetRequest() Request
	GetResponse() Response
}

type CommonResp map[string]interface{}

func (resp *CommonResp) Unmarshal(bts []byte) error {
	return json.Unmarshal(bts, resp)
}

type ErrorResponse struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	RequestID string `json:"request_id"`
	SubCode   string `json:"sub_code"`
	SubMsg    string `json:"sub_msg"`
}

package api

import "encoding/json"

type TpwdCreateReq struct {
	Url string
}

func (req *TpwdCreateReq) Params() map[string]string {
	return map[string]string{
		"url": req.Url,
	}
}

type TpwdCreateResp struct {
	ErrorResponse         ErrorResponse `json:"error_response"`
	TbkTpwdCreateResponse struct {
		Data struct {
			Model          string `json:"model"`
			PasswordSimple string `json:"password_simple"`
		} `json:"data"`
		RequestID string `json:"request_id"`
	} `json:"tbk_tpwd_create_response"`
}

func (resp *TpwdCreateResp) Unmarshal(bts []byte) error {
	return json.Unmarshal(bts, resp)
}

type TpwdCreateAPI struct {
	Request  *TpwdCreateReq
	Response *TpwdCreateResp
}

func NewTpwdCreateAPI() *TpwdCreateAPI {
	return &TpwdCreateAPI{
		Request:  new(TpwdCreateReq),
		Response: new(TpwdCreateResp),
	}
}

func (api *TpwdCreateAPI) Name() string {
	return "taobao.tbk.tpwd.create"
}

func (api *TpwdCreateAPI) GetRequest() Request {
	return api.Request
}

func (api *TpwdCreateAPI) GetResponse() Response {
	return api.Response
}

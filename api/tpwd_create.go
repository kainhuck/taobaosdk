package api

type TpwdCreateReq struct {
	Url string
}

func (req *TpwdCreateReq) Params() map[string]string {
	return map[string]string{
		"url": req.Url,
	}
}

type TpwdCreateAPI struct {
	Request  *TpwdCreateReq
	Response *CommonResp
}

func NewTpwdCreateAPI() *TpwdCreateAPI {
	return &TpwdCreateAPI{
		Request:  new(TpwdCreateReq),
		Response: new(CommonResp),
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

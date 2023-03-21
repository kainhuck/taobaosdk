package api

import (
	"strconv"
)

type ItemInfoReq struct {
	NumIids  string
	Platform int
	Ip       string
}

func (req *ItemInfoReq) Params() map[string]string {
	return map[string]string{
		"num_iids": req.NumIids,
		"platform": strconv.Itoa(req.Platform),
		"ip":       req.Ip,
	}
}

type ItemInfoAPI struct {
	Request  *ItemInfoReq
	Response *CommonResp
}

func NewItemInfoAPI() *ItemInfoAPI {

	return &ItemInfoAPI{
		Request:  new(ItemInfoReq),
		Response: new(CommonResp),
	}
}

func (api *ItemInfoAPI) Name() string {
	return "taobao.tbk.item.info.get"
}

func (api *ItemInfoAPI) GetRequest() Request {
	return api.Request
}

func (api *ItemInfoAPI) GetResponse() Response {
	return api.Response
}

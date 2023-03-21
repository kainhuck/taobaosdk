package api

import (
	"encoding/json"
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

type ItemInfoResp map[string]interface{}

func (resp *ItemInfoResp) Unmarshal(bts []byte) error {
	return json.Unmarshal(bts, resp)
}

type ItemInfoAPI struct {
	Request  *ItemInfoReq
	Response *ItemInfoResp
}

func NewItemInfoAPI() *ItemInfoAPI {

	return &ItemInfoAPI{
		Request:  new(ItemInfoReq),
		Response: new(ItemInfoResp),
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

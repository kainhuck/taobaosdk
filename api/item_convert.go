package api

import "strconv"

type ItemConvertReq struct {
	AdzoneID int
	Fields   string
	NumIIDS  string
}

func (req *ItemConvertReq) Params() map[string]string {
	return map[string]string{
		"adzone_id": strconv.Itoa(req.AdzoneID),
		"fields":    req.Fields,
		"num_iids":  req.NumIIDS,
	}
}

type ItemConvertAPI struct {
	Request  *ItemConvertReq
	Response *CommonResp
}

func NewItemConvertAPI() *ItemConvertAPI {
	return &ItemConvertAPI{
		Request:  new(ItemConvertReq),
		Response: new(CommonResp),
	}
}

func (api *ItemConvertAPI) Name() string {
	return "taobao.tbk.item.convert"
}

func (api *ItemConvertAPI) GetRequest() Request {
	return api.Request
}

func (api *ItemConvertAPI) GetResponse() Response {
	return api.Response
}

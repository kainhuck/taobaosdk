package taobaosdk

import "net/http"

type Option struct {
	HttpClient *http.Client
	BaseURI    string
}

var DefaultOption = &Option{
	HttpClient: http.DefaultClient,
	BaseURI:    FormalURI,
}

func (op *Option) Apply(op2 *Option) {
	if op2.HttpClient != nil {
		op.HttpClient = op2.HttpClient
	}

	if op2.BaseURI != "" {
		op.BaseURI = op2.BaseURI
	}
}

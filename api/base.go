package api

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

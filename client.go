package taobaosdk

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/kainhuck/taobaosdk/api"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Client struct {
	appKey     string
	appSecret  string
	baseURI    string
	httpClient *http.Client
}

func NewClient(appKey string, appSecret string, options ...*Option) *Client {
	option := DefaultOption
	for _, op := range options {
		option.Apply(op)
	}

	return &Client{
		appKey:     appKey,
		appSecret:  appSecret,
		baseURI:    option.BaseURI,
		httpClient: option.HttpClient,
	}
}

func (c *Client) Do(api api.Api) error {
	params := make(url.Values)
	// 公共参数
	params.Add("method", api.Name())
	params.Add("app_key", c.appKey)
	params.Add("timestamp", time.Now().Format("2006-01-02 15:04:05"))
	params.Add("v", "2.0")
	params.Add("format", "json")
	params.Add("sign_method", "md5")
	// 业务参数
	for k, v := range api.GetRequest().Params() {
		params.Add(k, v)
	}
	// 签名
	params.Add("sign", c.sign(params))

	bts, err := c.sendRequest(params)
	if err != nil {
		return err
	}

	return api.GetResponse().Unmarshal(bts)
}

// 签名
func (c *Client) sign(params url.Values) (str string) {
	var keyArr []string
	for key, _ := range params {
		keyArr = append(keyArr, key)
	}
	sort.Strings(keyArr)

	for _, key := range keyArr {
		str += key + params.Get(key)
	}

	s := md5.New()
	s.Write([]byte(c.appSecret + str + c.appSecret))
	return strings.ToUpper(hex.EncodeToString(s.Sum(nil)))
}

// 发送请求
func (c *Client) sendRequest(params url.Values) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, c.baseURI, strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

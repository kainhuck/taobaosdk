package main

import (
	"encoding/json"
	"fmt"
	"github.com/kainhuck/taobaosdk"
	"github.com/kainhuck/taobaosdk/api"
	"log"
	"os"
)

var (
	appKey    = os.Getenv("top_app_key")
	appSecret = os.Getenv("top_app_secret")
	client    = taobaosdk.NewClient(appKey, appSecret)
)

func main() {
	tpwdCreate := api.NewTpwdCreateAPI()
	tpwdCreate.Request.Url = "https://s.click.taobao.com/bKX3JKu"
	do(tpwdCreate)
}

func do(api api.Api) {
	if err := client.Do(api); err != nil { // 通过client发送请求
		log.Fatal(err)
	}
	bts, _ := json.MarshalIndent(api.GetResponse(), "", "  ")
	fmt.Println(string(bts)) // 获取返回信息
}

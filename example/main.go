package main

import (
	"fmt"
	"github.com/kainhuck/taobaosdk"
	"github.com/kainhuck/taobaosdk/api"
	"log"
	"os"
)

func main() {
	appKey := os.Getenv("top_app_key")
	appSecret := os.Getenv("top_app_secret")

	// 新建client
	client := taobaosdk.NewClient(appKey, appSecret)

	// 淘宝客-公用-淘宝客商品详情查询(简版) | 所有支持的接口都位于 github.com/kainhuck/taobaosdk/api 下
	itemInfo := api.NewItemInfoAPI()
	itemInfo.Request.NumIids = "123,456"        // 设置参数
	if err := client.Do(itemInfo); err != nil { // 通过client发送请求
		log.Fatal(err)
	}
	fmt.Println(itemInfo.Response) // 获取返回信息
}

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

	client := taobaosdk.NewClient(appKey, appSecret)

	itemInfo := api.NewItemInfoAPI()

	itemInfo.Request.NumIids = "123,456"

	err := client.Do(itemInfo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(itemInfo.Response)
}

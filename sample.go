package main

import (
	"fmt"
	"os"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/platform"
)

func main() {
	// create pixelbin config object
	config := platform.NewPixelbinConfig(
		"API_TOKEN",
		"https://api.pixelbin.io",
	)
	// set oauthclient
	config.SetOAuthClient()

	// create pixelbin client object
	pixelbin := platform.NewPixelbinClient(config)

	file, _ := os.Open("./tests/1.jpeg")

	// Parameters for FileUpload function
	params := platform.FileUploadXQuery{
		File: file,
	}
	result, err := pixelbin.Assets.FileUpload(params)

	if err != nil {
		fmt.Println(err)
	}
	// use result
	fmt.Println(result)
}

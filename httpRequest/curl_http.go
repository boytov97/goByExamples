package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type response2 struct {
	PostType     string `json:"type"`
	ThumbnailUrl string `json:"thumbnail_url"`
	StatusMsg    string `json:"status_msg"`
}

func main() {

	url := "https://www.tiktok.com/oembed?url=https%3A%2F%2Fwww.tiktok.com%2F%40myktybek027%2Fvideo%2F7031122951571377409%3Fis_copy_url%3D0&is_from_webapp=v1&sender_device=pc&sender_web_id=6914188754732123654"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body) // body type is []uint8
	//tiktokPostData := string(body)
	//tiktokPostSlice := []byte(tiktokPostData)

	res2 := response2{}
	json.Unmarshal([]byte(body), &res2)

	fmt.Println(res2)
	fmt.Println("sm:", res2.StatusMsg)

	if res2.StatusMsg == "" {
		fmt.Println("t:")
	}
}

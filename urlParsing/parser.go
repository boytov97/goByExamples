package main

import (
	"fmt"
	"net/url"
	"strconv"
)

func GetTiktokExpirationDateInUrl(rawUrl string) int64 {
	u, err := url.Parse(rawUrl)

	fmt.Println("err", err)

	if err != nil {
		return 0
	}

	q, err := url.ParseQuery(u.RawQuery)

	fmt.Println("err1", err)

	if err != nil {
		return 0
	}

	if !q.Has("x-expires") {
		fmt.Println("has not")
		return 0
	}

	expirationDateString := q.Get("x-expires")

	//fmt.Println("date:", reflect.TypeOf(expirationDateString))
	//expirationDateInt, err := strconv.Atoi(expirationDateString)

	expirationDateInt, err := strconv.ParseInt(expirationDateString, 16, 64)

	if err != nil {
		fmt.Println("er2", err)
		return 0
	}

	return expirationDateInt
}

func main() {
	parsed := GetTiktokExpirationDateInUrl("https://p16-sign-va.tiktokcdn.com/obj/tos-maliva-p-0068/937d1ef6b8394d7099342a8c77a7cc6e?x-expires=1637848800&x-signature=6A62m2S5fLHUQ2MLxIiq9kBGHQ8%3D")
	
	// test commit
	fmt.Println("parsed:", parsed)
}


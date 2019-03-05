package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("a", "aa")
	v.Add("b", "bb vv")
	v.Add("c", "cc 中国人")
	body := v.Encode()
	fmt.Println(body)
}

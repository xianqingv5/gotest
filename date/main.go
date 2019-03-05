package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Format("20060102"))
	date := time.Now().Format("200601021504")
	fmt.Println(date)
}

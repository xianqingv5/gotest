package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function mian.")
	panic(errors.New("someting wrong"))
	p := recover()
	fmt.Printf("panic: %s\n", p)
	fmt.Println("Exit function main.")
}

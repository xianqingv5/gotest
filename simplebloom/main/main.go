package main

import (
	"fmt"

	"github.com/willf/bitset"
)

func main() {
	var b bitset.BitSet
	b.Set(10).Set(11).Set(15)
	if b.Test(1000) {
		b.Clear(1000)
	}

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}

	fmt.Println(b.Len())
	fmt.Println(b.String())

	// var a uint
	// var d uint32
	// var c uint64
	// fmt.Println(unsafe.Sizeof(a))
	// fmt.Println(unsafe.Sizeof(d))
	// fmt.Println(unsafe.Sizeof(c))
}

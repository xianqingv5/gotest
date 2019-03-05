// https://www.colabug.com/643505.html

package main

import (
	"fmt"
	"math/rand"

	"github.com/willf/bitset"
)

func main() {
	fmt.Printf("Hello from BitSet!\n")
	var b bitset.BitSet

	for i := 0; i < 100; i++ {
		card1 := uint(rand.Intn(52)) // 生成随机数
		card2 := uint(rand.Intn(52))
		b.Set(card1) // Set bit card1 to 1
		// b.Set(card2)
		if b.Test(card2) { // 测试card2位是否设置。
			fmt.Println("Go Fish!")
		}
		b.Clear(card1) // Clear bit card1 to 0
	}

	b.Set(10).Set(11)

	// fmt.Println(uint(rand.Intn(52)))

	for i, e := b.NextSet(0); e; i, e = b.NextSet(i + 1) {
		fmt.Println("The following bit is set:", i)
	}

	if b.Intersection(bitset.New(100).Set(11)).Count() == 1 { // b sets 与 new sets的交集
		fmt.Println("Intersection works.")
	} else {
		fmt.Println("Intersection doesn't work???")
	}

	var d bitset.BitSet
	fmt.Println(d.Bytes())
	d.Set(0)
	fmt.Println(d.Bytes(), 0)
	fmt.Println()
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// var a int64 = 3
	// var b float64 = float64(a)
	// // a = float64(a)
	// fmt.Println(&a)
	// fmt.Println(&b)

	var n int64 = 3
	var pn = &n
	fmt.Println(pn)
	var pf = (*float64)(unsafe.Pointer(pn))
	fmt.Println(*pf)

	a := [4]int{0, 1, 2, 3}
	p1 := unsafe.Pointer(&a[1])                               // index为1的元素
	p3 := unsafe.Pointer(uintptr(p1) + 2*unsafe.Sizeof(a[3])) // 拿到index为3的指针
	*(*int)(p3) = 4                                           // 重新赋值
	fmt.Println(a)                                            // a = [0 1 2 4]
}

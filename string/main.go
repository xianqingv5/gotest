// GOLANG 格式化对象 STRING()
// http://vearne.cc/archives/344
package main

import "fmt"

// Car car
type Car struct {
	age  int
	name string
}

// Car2 car2
type Car2 struct {
	age  int
	name string
}

func (c Car) String() string {
	return fmt.Sprintf("Car-[name:%v, age:%v]", c.name, c.age)
}

func main() {
	c1 := Car{age: 2, name: "buick"}
	c2 := Car{age: 1, name: "benz"}
	fmt.Println(c1)

	// 对于切片中的对象，也能递归调用String()
	carSlice := []Car{}
	carSlice = append(carSlice, c1)
	carSlice = append(carSlice, c2)
	fmt.Println(carSlice)

	c3 := Car2{age: 10, name: "songhq"}
	fmt.Println(c3)
}

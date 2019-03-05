package main

import (
	"fmt"
	"os"
	"os/signal"
)

func signalListen() {
	c := make(chan os.Signal)
	signal.Notify(c)
	for {
		s := <-c

		// 收到信号后的处理，这里只是输出信号内容，可以做一些更有意思的事
		fmt.Println("get signal:", s)
		os.Exit(1)
	}
}

func main() {
	ch := make(chan os.Signal)

	// 如果不指定要监听的信号，那么默认是所有信号
	signal.Notify(ch)

	// 停止向ch转发信号，ch将不再收到任何信号
	signal.Stop(ch)

	// ch将一直阻塞在这里，因为它将收不到任何信号
	// 所以下面的exit输出也无法执行
	<-ch
	fmt.Println("exit")
}

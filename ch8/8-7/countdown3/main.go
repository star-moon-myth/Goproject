package main

import "fmt"

func main() {
	ch := make(chan int, 2) // 创建一个 buffer 大小为 1 的 int channel

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // 从 ch 接收数据
			fmt.Println(x) // 打印接收到的数据
		case ch <- i: // 向 ch 发送数据 i
			fmt.Println("sent", i) // 打印发送的数据
		}
	}

}
package main

import "fmt"

func main() {
	natural := make(chan int)
	square := make(chan int)
	go counter(natural)
	go squarer(natural, square)
	print(square)
}
func counter(out chan int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
}
func squarer(in chan int, out chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func print(in chan int) {
	for i := range in {
		fmt.Println(i)
	}
}
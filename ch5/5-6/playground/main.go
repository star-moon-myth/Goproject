package main

import "fmt"

func greet(prefix string, who ...string) {
	if len(who) == 0 {
		fmt.Println(prefix + "nobody")
		return
	}

	for _, name := range who {
		fmt.Println(prefix + name)
	}
}

func main() {
	names := []string{"Alice", "Bob", "Charlie"}

	greet("Hello, ", names...) // 使用 ... 展开切片

	greet("Goodbye, ")
}

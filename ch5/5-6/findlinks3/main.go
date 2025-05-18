package main

import (
	"fmt"
	"log"
	"os"
)

func extract(url string) ([]string, error) {
	switch url {
	case "https://golang.org":
		return []string{
			"https://go.dev",
			"https://pkg.go.dev",
			"https://play.golang.org",
		}, nil
	case "https://go.dev":
		return []string{
			"https://golang.org",
			"https://blog.golang.org",
		}, nil
	default:
		return []string{}, nil
	}
}
func breadfirst(f func(link string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist)>0 {
		links := worklist
		worklist = nil
		for _, link := range links {
			if !seen[link] {
				seen[link] = true
				worklist = append(worklist, f(link)...)
			}
		}
	}
}
func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url) // 使用模拟的 extract 函数
	if err != nil {
		log.Print(err)
	}
	return list
}
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadfirst(crawl, os.Args[1:])
}
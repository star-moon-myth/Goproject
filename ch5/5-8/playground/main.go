package main

import (
	"fmt"
	"net/http"
	"strings"
	"golang.org/x/net/html"
	"log"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err // 如果获取网页失败，直接返回错误
	}
	defer resp.Body.Close() // 确保 resp.Body 被关闭，防止资源泄漏

	// 检查 Content-Type 是否为 HTML (例如 "text/html; charset=utf-8")
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct) // 返回详细的错误信息
	}

	doc, err := html.Parse(resp.Body)
	//resp.Body.Close() // resp.Body 已经在上面用 defer 关闭了，这里不需要重复关闭
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err) // 返回详细的错误信息
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data) // 打印 title 标签的内容
		}
	}

	forEachNode(doc, visitNode, nil) // 使用 forEachNode 遍历 HTML 树
	return nil // 如果一切正常，返回 nil 表示没有错误
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func main() {
	url := "https://www.example.com"
	if err := title(url); err != nil {
		log.Fatal(err) // 如果 title 函数返回错误，打印错误并终止程序
	}
}

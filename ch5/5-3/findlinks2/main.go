package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html" // 使用官方的 HTML 解析库
)

func main() {
	// 遍历命令行参数，从索引 1 开始（跳过程序名）
	for _, url := range os.Args[1:] {
		// 调用 findLinks 函数获取指定 URL 中的链接
		links, err := findLinks(url)
		// 如果 findLinks 函数返回错误（例如网络错误、HTTP 错误）
		if err != nil {
			// 将错误信息输出到标准错误流
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			// 继续处理下一个 URL
			continue
		}
		// 遍历 findLinks 函数返回的链接列表
		for _, link := range links {
			// 将每个链接输出到标准输出流
			fmt.Println(link)
		}
	}
}

// findLinks 函数：发起 HTTP 请求，解析 HTML，提取链接
func findLinks(url string) ([]string, error) {
	// 发起 HTTP GET 请求
	resp, err := http.Get(url)
	// 如果请求失败（例如网络错误）
	if err != nil {
		return nil, err // 直接返回错误
	}

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		// 关闭响应体
		resp.Body.Close()
		// 返回一个包含状态码的错误
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	// 解析 HTML
	doc, err := html.Parse(resp.Body)
	// 关闭响应体（无论解析是否成功，都要关闭）
	resp.Body.Close()
	// 如果解析 HTML 失败
	if err != nil {
		// 返回一个包含解析错误的错误
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	// 调用 visit 函数提取链接
	return visit(nil, doc), nil
}

// visit 函数：递归遍历 HTML 节点树，提取链接
func visit(links []string, n *html.Node) []string {
	// 如果当前节点是元素节点，且标签名为 "a"（超链接）
	if n.Type == html.ElementNode && n.Data == "a" {
		// 遍历节点的所有属性
		for _, a := range n.Attr {
			// 如果属性名为 "href"（链接地址）
			if a.Key == "href" {
				// 将链接地址添加到 links 切片中
				links = append(links, a.Val)
			}
		}
	}

	// 递归遍历当前节点的子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// 递归调用 visit 函数，并将 links 切片传递下去
		links = visit(links, c)
	}

	// 返回更新后的 links 切片
	return links
}

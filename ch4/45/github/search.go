package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues queries the GitHub issue tracker.
// SearchIssues 函数查询 GitHub 的 issue 跟踪器。
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " ")) // 对查询条件进行 URL 编码
	resp, err := http.Get(IssuesURL + "?q=" + q)    // 发起 HTTP GET 请求
	if err != nil {
		return nil, err // 如果请求失败，返回错误
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	// 我们必须在所有执行路径上关闭 resp.Body。
	// （第 5 章介绍了 'defer'，这使得关闭操作更简单。）
	if resp.StatusCode != http.StatusOK { // 检查响应状态码是否为 200 OK
		resp.Body.Close()                                 // 如果不是，关闭响应体
		return nil, fmt.Errorf("search query failed: %s", resp.Status) // 返回错误信息
	}

	var result IssuesSearchResult                         // 定义一个 IssuesSearchResult 类型的变量来存储解码后的结果
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // 将响应体解码为 IssuesSearchResult 结构
		resp.Body.Close() // 如果解码失败，关闭响应体
		return nil, err // 返回错误
	}
	resp.Body.Close() // 解码成功，关闭响应体
	return &result, nil // 返回解码后的结果和 nil 错误
}

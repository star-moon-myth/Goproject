package main

import (
	"encoding/json"
	"fmt"
	"html/template" // 导入 html/template 包，用于生成 HTML
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// 定义 GitHub API 的 URL
const IssuesURL = "https://api.github.com/search/issues"

// 定义一个结构体来表示 GitHub issue 的数据
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"` // issue 的总数
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"` // issue 页面的 URL
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"` // issue 创建时间
	Body      string                        // issue 的正文内容
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"` // 用户页面的 URL
}

// 定义一个 HTML 模板，用于生成 issue 列表
// template.Must 用于包装模板解析，并在发生错误时 panic
var issueList = template.Must(template.New("issueList").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Items}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

// searchIssues 函数向 GitHub API 发送请求并获取 issue 列表
func searchIssues(terms []string) (*IssuesSearchResult, error) {
	// 构建查询参数
	q := url.QueryEscape(strings.Join(terms, " ")) // 将搜索条件连接成字符串，并进行 URL 编码
	// 发送 GET 请求
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 确保在函数返回时关闭响应体
	defer resp.Body.Close()

	// 检查 HTTP 状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// 解析 JSON 数据
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func main() {
	// 从命令行参数中获取搜索条件
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// 执行 HTML 模板，并将结果输出到标准输出
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

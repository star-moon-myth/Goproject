package main

import (
	"fmt"
	"log"
	"strings"
	"golang.org/x/net/html"
)
var depth int
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		// 打印缩进和开始标签
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		// 增加深度
		depth++
	} else if n.Type == html.TextNode {
		// 处理文本节点，去除多余的空白
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%*s%s\n", depth*2, "", text)
		}
	}
}

func endElement(n *html.Node){
	if n.Type==html.ElementNode{
		depth--
		fmt.Printf("%*s<%s>\n",depth*2,"",n.Data)
	}
}
func foreachNode(n *html.Node,pre,post func(n *html.Node)){
	if pre!=nil{
		pre(n)
	}
	for c:=n.FirstChild;c!=nil;c=c.NextSibling{
		foreachNode(c,pre,post)
	}
	if post!=nil{
		post(n)
	}
}
func main(){
	doc, err := html.Parse(strings.NewReader(`
		<html>
		<head>
		    <title>我的网页</title>
		</head>
		<body>
		    <h1>欢迎</h1>
		    <p>这是一个段落。</p>
		    <div>
		        <img src="image.jpg" alt="图片">
		    </div>
		</body>
		</html>
	`))
	if err!=nil{
		log.Fatal(err)
	}
	foreachNode(doc,startElement,endElement)
}
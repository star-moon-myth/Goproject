package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(inttostring([]int{1,3,3}))

}
func inttostring(values []int) string{
	// 建立缓冲区
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i,v:=range values{
		if i>0{
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf,"%d",v)				
	}
	buf.WriteByte(']')	
	return buf.String()
}

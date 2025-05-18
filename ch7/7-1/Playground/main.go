package main

import "fmt"

type count int

func (c *count) Write(b []byte) (int, error) {
	*c += count(len(b))
	return len(b), nil
}
func main() {
	var co count
	co.Write([]byte("delly"))
	fmt.Println(co)
    co=0
    var name ="delly"
    fmt.Fprintf(&co,"haha%s",name)
    fmt.Println(co)
    

}
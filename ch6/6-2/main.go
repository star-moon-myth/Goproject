package main

import (
	"fmt"
	"net/url"
)

func main() {
	// s := url.Values{"lang":{"aaa","eee"}}
	// url.Values{"lang":{"aaa","eee"}}.Add("lala对","🦐")
	// s.Add("lala对","🦐")
	// s.Add("hobby","eat")
	// s.Add("hobby","drink")
	// s.Add("hobby","play")
	// s.Add("age","34")	
	// s.Add("lang","ccc")	
	// s.Set("hobby","😭")
	// fmt.Println(s.Get("age"))
	// fmt.Println(s.Get("hobby"))
	// fmt.Println(s.Get("lang"))
	// fmt.Println(s)
	s:=url.Values{}
	s.Add("lala对","🦐")
	fmt.Println(s)
	s=nil
	s.Add("lala对","🦐")
	fmt.Println(s)


}
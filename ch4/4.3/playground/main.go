package main

import (
	"fmt"
	"sort"
)

func main() {
	// ages:=make(map[string]int)
	ages := map[string]int{
		"mike": 23,
		"jina": 34,
		"lina": 27,
	}
	ages["alice"] = 89
	delete(ages,"alice")
	ages["nike"]++
	// for key,value:=range ages{
	// 	fmt.Println(key,value)
	// }
	// var s []string
	s:=make([]string,0,len(ages))

	for name:=range ages{
		s=append(s, name)
	}
	fmt.Printf("%#v\n",s)
	sort.Strings(s)
	fmt.Printf("%#v\n",s)
	fmt.Println(s)
	for _,name :=range s{
		fmt.Println(name ,ages[name])
	}
	
}
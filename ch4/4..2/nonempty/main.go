package main

import (
	"fmt"
)
// func nonempty(strings []string)[]string{
// 	i:=0
// 	for _,value:=range strings{
// 		if value!=""{
// 			strings[i]=value
// 			i++
// 		}
// 	}
// 	return strings[:i]
// }
func remove(strings []int,i int)[]int{
	// copy(strings[i:],strings[i+1:])
	strings[i]=strings[len(strings)-1]
	return strings
}

func main(){
	// data:=[]string{"q","b","","n","","m"}
	// data = nonempty(data)
	// fmt.Printf("%#v\n", data)	
	// fmt.Println(data)
	s:=[]int{1,2,3,4,5,5,6,7,8,9}
	fmt.Println(remove(s,2))
}
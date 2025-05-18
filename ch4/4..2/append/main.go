package main

import "fmt"

func main() {
	fmt.Println(appendint([]int{1,2,3,4,5,6,7},9))
}
func appendint(x []int,y int)[]int{
	var z []int
	zlen:=len(x)+1
	if zlen<=cap(x){
		z=x[:zlen]
		z[len(x)]=y
	}else{
		zcap:=zlen
		if zcap<2*len(x){
			zcap=2*len(x)
		}

		z=make([]int,zlen,zcap)
		copy(z,x)
		z[len(x)]=y
	}
	return z
}
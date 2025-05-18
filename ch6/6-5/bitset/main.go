package main

import (
	"bytes"
	"fmt"
)

type intset struct {
	words []uint64
}

func (i *intset) has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(i.words) && i.words[word]&(1<<bit) != 0
}
func (i *intset) add(x int) {
	word, bit := x/64, uint(x%64)
	for word >=len(i.words) {
		i.words = append(i.words, 0)
	}
	i.words[word] |= (1 << bit)
}
func (i *intset) union(n *intset) {
	for j, k := range n.words {
		if j < len(i.words) {
			i.words[j] = k
		} else {
			i.words = append(i.words, k)
		}
	}
}
func (s *intset) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for m,n:=range s.words{
		if n==0{
			continue
		}
		for j:=0;j<64;j++{
			if n&(1<<j)!=0{
				if buf.Len()>len("{"){
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf,"%d",64*m+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func main(){
	var x,y intset
	x.add(5)
	x.add(55)
	x.add(656)
	fmt.Println(x.String())


	y.add(9)
	y.add(42)
	fmt.Println(y.String()) 

	x.union(&y)
	fmt.Println(x.String())
	fmt.Println(x.has(9),x.has(89))


	fmt.Println(&x)         
	fmt.Println(x.String())    
	fmt.Printf("%x\n", x)

}
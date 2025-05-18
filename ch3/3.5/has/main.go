package main
import "fmt"
func Hadprefex(m,n string) bool {
	return len(m)>=len(n) && m[:len(n)]==n
}
func Hassufix(m,n string)  bool{
	return len(m)>=len(n) && m[len(m)-len(n):]==n
}
func Contain(m,n string) bool{
	for i:=0;i<len(m);i++{
		if Hadprefex(m[i:],n){
			return true
		}
	}
	return false
}
func main(){
	 m,n:="chinahahalove","love"
	 fmt.Println(Contain(m,n))
	 fmt.Println(Hadprefex(m,n))
	 fmt.Println(Hassufix(m,n))

}
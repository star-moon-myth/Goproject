package main
import "fmt"
type C float64
type F float64
const (
	Absolutezero C=-273.15
	Freezing C=0
	boil C =100
)
const(
	Absolutezerof F=-459.66999999
	Freezingf F=32
	boilf F=212
)
func Ftoc(c C)F{
	return F(c*9/5 + 32)
}
func Ctof(f F)C{
	return C((f - 32) * 5 / 9)
}
func main(){
	fmt.Println(Ftoc(Absolutezero))
	fmt.Println(Ftoc(Freezing))
	fmt.Println(Ftoc(boil))
	fmt.Println("--------------分线-------------")
	fmt.Println(Ctof(boilf))
	fmt.Println(Ctof(Absolutezerof))
	fmt.Println(Ctof(Freezingf))
	
}
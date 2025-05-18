package main
import "fmt"
func main() {
	const f,c=212,32
	fmt.Println(ftoc(f),ftoc(c))
}
func ftoc(c float32) float32 {
	return (c  - 32) * 5 / 9
}
package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct{ Celsius }


func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}


//String和Set实际上是在实现flag的value接口，CelsiusFlag类型结构体实现set，同时又继承Celsius的String方法
//至此celsiusFlag实现了接口，flag.CommandLine.Var(&f, name, usage)中f必须是一个接口类型
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }


func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) 
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//原来应该是flag.CommandLine.Var(&f, name, usage)这种，如今用函数CelsiusFlag进行封装，
//在自定义类型的情况下使赋值方式和原来一样（相当于变相的flag.Var，这样应该也可以），返回的是*Celsius（指向Celsius）指针
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

//常规命令行参数解析步骤

// func main(){
// 	// var name string
// 	// flag.StringVar(&name,"name","nina","请输入你的名字：")
// 	//或者
// 	name:=flag.String("name","nina","请输入你的名字：")
// 	flag.Parse()
// 	fmt.Println(name)
// }





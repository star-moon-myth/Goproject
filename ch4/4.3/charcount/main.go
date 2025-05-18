package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	countstring := make(map[rune]int)
	utfcount := [utf8.UTFMax+1]int{}
	invalid:=0
	input:=bufio.NewReader(os.Stdin)
	for {
		rune,n,err:=input.ReadRune()
		if err==io.EOF{
			break
		}
		if err != nil{
			fmt.Fprintf(os.Stderr,"%v\n",err)
			os.Exit(1)
		}
		if n==1 && rune==unicode.ReplacementChar{
			invalid++
			continue
		
		}
		
		utfcount[n]++
		countstring[rune]++

	}
	fmt.Printf("rune\tcount\n")		
	for r,c:=range countstring{
		fmt.Printf("%q\t%d\n",r,c)
	}
	fmt.Print("\nlen\tcount\n")
	for i,n:=range utfcount{
		if i>0{
			fmt.Printf("%d\t%d\n",i,n)
		}
	}
	if invalid>0{
		fmt.Println(invalid)
	}


}
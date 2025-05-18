package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type track struct {
	title  string
	artist string
	album  string
	year   int
	length time.Duration
}
type customsort struct{
	t []*track
	less func(x,y *track)bool
}

func le(s string)time.Duration{
	d,err:=time.ParseDuration(s)
	if err!=nil{
		panic(s)
	}
	return d 
}
func printformat(t []*track){
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw:=new(tabwriter.Writer).Init(os.Stdout,0,8,2,' ',0)
	fmt.Fprintf(tw,format,"Title","Artist", "Album", "Year","Length")
	fmt.Fprintf(tw,format,"-----", "------", "-----", "----", "------")
	for _,n:=range t{
		fmt.Fprintf(tw,format,n.title,n.artist, n.album, n.year,n.length)
	}
	tw.Flush()
	
}
type ks customsort
func(s ks)Len()int {
	return len(s.t)
}
func(s ks )Less(i,j int)bool{
	return s.less(s.t[i], s.t[j])
}
func(s ks)Swap(i,j int){
	s.t[i],s.t[j]=s.t[j],s.t[i]
}
func main(){
	tracks := []*track{
		{"Go", "Delilah", "From the Roots Up", 2012, le("3m38s")},
		{"Go", "Moby", "Moby", 1992, le("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, le("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, le("4m24s")},
	}
	fmt.Println("原来的凌乱的")
	printformat(tracks)
	fmt.Println("整理后的")
	custom:=ks{
		t: tracks,
		less:func(x,y *track)bool{
			if x.title!=y.title{
				return x.title<y.title
			}
			if x.year != y.year {   
				return x.year < y.year
			}
			if x.length != y.length { 
				return x.length < y.length
			}
			return false 
		} ,
	}
	fmt.Println("😭😭😭😭😭😭😭😭😭😭😭😭😭😭😭😭")
	sort.Sort(custom)
	printformat(tracks)
}

package main

import (
	"fmt"
	"sort"
)

// prereqs记录了每个课程的前置课程
var prereqs = map[string][]string{
    "algorithms": {"data structures"},
    "calculus": {"linear algebra"},
    "compilers": {
        "data structures",
        "formal languages",
        "computer organization",
    },
    "data structures":       {"discrete math"},
    "databases":             {"data structures"},
    "discrete math":         {"intro to programming"},
    "formal languages":      {"discrete math"},
    "networks":              {"operating systems"},
    "operating systems":     {"data structures", "computer organization"},
    "programming languages": {"data structures", "computer organization"},
}
func main(){
	for i,couses:=range toposort(prereqs){
		fmt.Printf("%d:\t%s\n",i+1,couses)
	}
}
func toposort(n map[string][]string)[]string{
	var seen=make(map[string]bool)
	// var s=map[string]bool{}
	var order []string
	var visitall func(items []string)
	visitall = func(items []string){
		for _,item:=range items{
			if !seen[item]{
				seen[item]=true
				visitall(n[item])
				order=append(order, item)
			}
		}
	}
	var keys []string
	for key :=range n{
		keys=append(keys,key)
	}
	sort.Strings(keys)
	visitall(keys)
	return  order 

}
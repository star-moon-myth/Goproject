package main

import "testing"

func TestAdd(t *testing.T){
	a:=1
	b:=2
	predict:=3
	if Add(a,b)!=predict{
		t.Errorf("anser is %d,but is %d",predict,Add(a,b))
	}
}
package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("resourse not found")

func findSourse(id string)error {
	if id=="not_exist"{
		return errNotFound
	}
	return nil
}

func proscess(user string)error{
	err:=findSourse(user)
	if err!=nil{
		return fmt.Errorf("process error:%w",errNotFound)
	}
	fmt.Println("sucess!🌸")
	return nil
}

func main(){
	err:=proscess("not_exist")
	if err!=nil{
		fmt.Printf("top error%v\n",err)
		if errors.Is(err,errNotFound){
			fmt.Println("the underlying cause was :resourse not found.")
		}
	}
}
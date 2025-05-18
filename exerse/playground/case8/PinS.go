package main

import (
	"encoding/json"
	"fmt"
)

type Userprofile struct {
	Name     string  `json:"username"`
	Nickname *string `json:"nickname,omitempty"`
	Age      *int    `json:"age,omitempty"`
}

func main() {
	user1 := Userprofile{Name: "bob"}
	nick := "haha"
	age := 23
	user1.Age = &age
	user1.Nickname = &nick
	s,_:=json.MarshalIndent(user1,"","  ")
	fmt.Println("\nuser1:")
	fmt.Println(string(s))

	user2:=Userprofile{Name:"nina"}
	s,_=json.MarshalIndent(user2,"","  ")
	fmt.Println("\nuser2:")
	fmt.Println(string(s))

	jsondata:=`{"username":"alice","nickname":null,"age":34}`
	var user3 Userprofile
	json.Unmarshal([]byte(jsondata),&user3)
	fmt.Println("\nuser3:")
	fmt.Println(user3.Name)
	if user3.Age!=nil{
		fmt.Println(*user3.Age)
	}else{
		fmt.Println("error:it is null")

	}
	if user3.Nickname!=nil {
		fmt.Println(*user3.Nickname)
	}else{
		fmt.Println("error:it is null")
	}

}
package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))
	fmt.Println(reflect.TypeOf(bs)) // []uint8
	fmt.Println(string(bs))

	loginPword1 := `password1234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")
}

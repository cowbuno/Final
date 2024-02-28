package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	lox1, _ := bcrypt.GenerateFromPassword([]byte("lox"), bcrypt.DefaultCost)

	fmt.Println(string(lox1))
	fmt.Println(bcrypt.CompareHashAndPassword(lox1, []byte("lox")))
}

package main

import (
	"flag"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

var flagLaravel bool
var flagPassword string

func main() {
	flag.BoolVar(&flagLaravel, "laravel", false, "Use Laravel helper")
	flag.StringVar(&flagPassword, "password", "", "this password")
	flag.Parse()

	if flagLaravel {
		laravelFunc()
	}
}

func laravelFunc() {
	if flagPassword != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(flagPassword), 10)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println("password:", flagPassword)
		fmt.Println("hash:", string(bytes))
	}
}

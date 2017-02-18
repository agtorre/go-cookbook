package main

import "github.com/agtorre/go-cookbook/chapter5/database"

func main() {
	if err := database.Exec(); err != nil {
		panic(err)
	}
}

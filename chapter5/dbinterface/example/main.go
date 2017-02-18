package main

import "github.com/agtorre/go-cookbook/chapter5/dbinterface"

func main() {
	if err := dbinterface.Exec(); err != nil {
		panic(err)
	}
}

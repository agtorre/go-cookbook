package basicerrors

import (
	"errors"
	"fmt"
)

// ErrorTyped is a way to make a package level
// error to check against. I.e. if err == TypedError
var ErrorTyped = errors.New("this is a typed error")

//BasicErrors demonstrates some ways to create errors
func BasicErrors() {
	err := errors.New("this is a quick and easy way to create an error")
	fmt.Println("errors.New: ", err)

	err = fmt.Errorf("an error occurred: %s", "something")
	fmt.Println("fmt.Errorf: ", err)

	err = ErrorTyped
	fmt.Println("typed error: ", err)
}

package greet

import "fmt"

// Hello prints a greeting to the specified recipient.
//
//go:fix inline
func Hello(to string) {
	fmt.Printf("Hello, %s!\n", to)
}

/*
This package provides a simple example of how to use Testable Examples.

https://go.dev/blog/examples

The name of file that contains examples must end with "_test.go".
Examples can be planced in the same file as normal tests, or they can be placed in a separate file.
As a convention it is common to name the separate file "example_test.go", but it can be any name as long as it ends with "_test.go",
If multiple example files are needed, it is common to name them "example_<xxx>_test.go".

The function name must start with "Example" and can be followed by an optional suffix that describes the example.
If suffix is used, it must match the name of the function or method being demonstrated, otherwise the example fails to compile.
*/
package example

import "fmt"

// Hello prints a greeting to the specified person.
func Hello(to string) {
	fmt.Printf("Hello, %s!\n", to)
}

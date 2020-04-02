package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

// Instructor:
/* Get more practice with interfaces implemented by the standard library.

Create a program that reads the contents of a text file then prints its contents to the terminal

The file to open should be provided as an argument to the program when it is executed at the terminal. For example,
'go run main.go myfile.txt' should open up the myfile.txt file

To read in the arguments provided to a program, you can reference the variable 'os.Args', which is a slice of type string

To open up a file, check out the documentation for the 'Open' function in the 'os' package - https://golang.org/pkg/os/

What interfaces does the 'File' type implement?

If the 'File' type implements the 'Reader' interface, you might be able to use reuse that io.Copy function
*/

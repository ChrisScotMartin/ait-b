package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-c/pkgmodc"
)

// ait-b imports pkg/mod from  ait-c
func main() {
	fmt.Println("Hello World")
	pkgmodc.Hello()
}

func SomeLibraryFunc() {
	fmt.Println("I'm a func from the main root module ait-a")
	pkgmodc.Hello()
}

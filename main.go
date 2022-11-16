package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-a/pkgmoda"
	"github.com/ChrisScotMartin/ait-c/pkgmodc"
)

// ait-b imports pkg/mod from  ait-c
func main() {
	fmt.Println("Hello World")
	pkgmodc.Hello()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-a")
}

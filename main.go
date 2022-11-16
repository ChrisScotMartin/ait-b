package main

import (
	"fmt"
	"github.com/ChrisScotMartin/ait-b/pkgmodb"
)

func main() {
	fmt.Println("Hello World")
	pkgmodb.Hello()
}

func SomeServerFunc() {
	fmt.Println("I'm a func from the main root module ait-b")
}

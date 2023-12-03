package main

import (
	"fmt"
	"github.com/leighmcculloch/go-strrev"
)

func main() {
	phrase := "Hello, OTUS!"
	fmt.Println(strrev.Reverse(phrase))
}

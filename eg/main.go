package main

import (
	"fmt"
	"github.com/TaigaMikami/pokesay"
)

func main() {
	simple()
}

func simple() {
	say, err := pokesay.Say(
		pokesay.Phrase("Hello\nWorld"),
		pokesay.Type("Pikachu"),
		pokesay.Thinking(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}

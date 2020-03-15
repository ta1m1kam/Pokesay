package main

import (
	"fmt"
	pokesay "github.com/TaigaMikami/Pokesay"
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

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
		pokesay.Phrase("Hello"),
		pokesay.Type("Bulbasaur"),
		pokesay.Thinking(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}

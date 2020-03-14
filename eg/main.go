package main

import (
	"flag"
	"fmt"
	"github.com/TaigaMikami/pokesay/ballon"
	"github.com/TaigaMikami/pokesay/img2xterm"
	"image"
	"os"
)


func main() {
	simple()
}

func simple() {
	cows := `         \
          \
`
	flag.Parse()
	inputs := ballon.ReadInput(flag.Args())
	width := ballon.MaxWidth(inputs)
	messages := ballon.SetPadding(inputs, width)
	balloon := ballon.ConstructBallon(messages, width)
	fmt.Println(balloon)
	fmt.Println(cows)
	file, _ := os.Open("images/Bulbasaur.png")
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	img2xterm.Img2xterm(img)
}



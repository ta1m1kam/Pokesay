package main

import (
	"bufio"
	"fmt"
	pokesay "github.com/TaigaMikami/Pokesay"
	"io"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "pokemon",
				Aliases: []string{"p"},
				Value: "",
				Usage: "language for the greeting",
			},
			&cli.BoolFlag{
				Name: "random",
				Aliases: []string{"r"},
				Value: true,
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			phrase := "Hello\nworld!!"
			if c.NArg() > 0 {
				phrase = c.Args().Get(0)
			} else if isPipe() {
				phrase = pipeText()
			}

			if c.String("pokemon") == "" {
				say, err := pokesay.Say(
					pokesay.Phrase(phrase),
					pokesay.Random(),
				)
				if err != nil {
					panic(err)
				}
				fmt.Print(say)
			} else {
				say, err := pokesay.Say(
					pokesay.Phrase(phrase),
					pokesay.Type(c.String("pokemon")),
				)
				if err != nil {
					panic(err)
				}
				fmt.Print(say)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func isPipe() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Size())
	return info.Size() > 0
}

func pipeText() string {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	return strings.TrimRight(string(output),"\r\n")
}

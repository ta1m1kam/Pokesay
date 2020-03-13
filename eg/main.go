package main

import (
	"fmt"
	"github.com/TaigaMikami/pokesay"
	"math"
)

func main() {
	simple()
}

func simple() {
	say, err := pokesay.Say(
			pokesay.Phrase("Hello"),
			pokesay.Type("Pikachu"),
		)
	if err != nil {
		panic(err)
	}
	fmt.Println(say)
}

func rawRGB2BrightnessPixels(raw []byte) (ret []float64) {
	for cur := 0; cur < len(raw); cur += 3 {
		r, g, b := raw[cur], raw[cur+1], raw[cur+2]
		bri := (float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114) / 255.0
		ret = append(ret, bri)
	}
	return
}

func floatMod(x, y float64) float64 {
	return x - y*math.Floor(x/y)
}

func floatMin(x, y float64) float64 {
	if x-y > 0 {
		return y
	}
	return x
}

func AsciiDrawRGB(raw []byte) {
	var chr string
	width, height := 100, 100
	pixels := rawRGB2BrightnessPixels(raw)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			brightness := pixels[y*width+x]
			bg := brightness*23 + 232
			fg := floatMin(255, bg+1)
			mod := floatMod(bg, 1.0)

			switch {
			case mod < 0.2:
				chr = " "
			case mod < 0.4:
				chr = "░"
			case mod < 0.6:
				chr = "▒"
			case mod < 0.8:
				bg, fg = fg, bg
				chr = "▒"
			default:
				bg, fg = fg, bg
				chr = "░"
			}

			fmt.Printf("\033[48;5;%dm\033[38;5;%dm%s", int(bg), int(fg), chr)
		}
		if (y + 1) < height {
			fmt.Print("\n")
		}
	}
}

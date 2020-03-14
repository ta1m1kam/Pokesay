package img2xterm

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"github.com/gookit/color"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func getPixels(img image.Image) ([][]Pixel) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}
	return pixels
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func Img2xterm(img image.Image) {
	pixels := getPixels(img)
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		if y < 10 {
			continue
		}
		for x := 0; x < width; x++ {
			if pixels[y][x].A == 0 {
				fmt.Print("　")
			} else {
				c := color.RGB(uint8(pixels[y][x].R), uint8(pixels[y][x].R), uint8(pixels[y][x].R), true)
				c.Print("　")
			}
		}
		fmt.Println()
	}
}


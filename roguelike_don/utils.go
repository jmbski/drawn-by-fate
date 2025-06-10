package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func CheckErr[T any](returnVal T, err error) T {
	if err != nil {
		log.Fatal(err)
	}
	return returnVal
}

func GetFontOpts(x, y int, fontColors ...color.Color) *text.DrawOptions {
	var fontColor color.Color = color.White
	if len(fontColors) > 0 {
		fontColor = fontColors[0]
	}

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	opts.ColorScale.ScaleWithColor(fontColor)

	return opts

}

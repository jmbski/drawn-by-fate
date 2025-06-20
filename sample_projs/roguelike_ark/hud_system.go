package main

import (
	"fmt"
	"image/color"
	"roguelike/fonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var hudImg *ebiten.Image
var hudFont *text.GoXFace

type HudWriter struct {
	Screen *ebiten.Image
	X      int
	Y      int
	Color  color.Color
}

func NewHudWriter(screen *ebiten.Image, x, y int) *HudWriter {
	return &HudWriter{
		Screen: screen,
		X:      x,
		Y:      y,
		Color:  color.White,
	}
}

func (o *HudWriter) fontOpts() *text.DrawOptions {
	return GetFontOpts(o.X, o.Y, o.Color)
}

func (o *HudWriter) write(msg string) {
	text.Draw(o.Screen, msg, hudFont, o.fontOpts())
	o.Y += YPad
}

func ProcessHUD(g *Game, screen *ebiten.Image) {
	if hudImg == nil {
		hudImg = GetImage("UIPanel.png")
	}

	if hudFont == nil {
		tt := CheckErr(opentype.Parse(fonts.MPlus1pRegular_ttf))

		const dpi = 72
		hudFont = text.NewGoXFace(CheckErr(opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    FontSize,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})))

	}

	gd := NewGameData()

	uiY := (gd.ScreenHeight - gd.UIHeight) * gd.TileHeight
	uiX := (gd.ScreenWidth * gd.TileWidth) / 2
	fontX := uiX + XPad
	fontY := uiY + YPad

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(uiX), float64(uiY))
	screen.DrawImage(hudImg, opts)

	hudWriter := NewHudWriter(screen, fontX, fontY)
	query := g.Context.PlayerHuds.Query()

	for query.Next() {
		_, h, ac, wpn := query.Get()

		hudWriter.write(fmt.Sprintf("Health: %d / %d", h.CurrentHealth, h.MaxHealth))

		hudWriter.write(fmt.Sprintf("Armor Class: %d", ac.ArmorClass))
		hudWriter.write(fmt.Sprintf("Defense: %d", ac.Defense))

		hudWriter.write(fmt.Sprintf("Damage: %d - %d", wpn.MinDamage, wpn.MaxDamage))
		hudWriter.write(fmt.Sprintf("To Hit Bonus: %d", wpn.ToHitBonus))

	}
}

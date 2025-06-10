package main

import (
	"image/color"
	"log"

	"github.com/RAshkettle/rrogue/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var userLogImg *ebiten.Image = nil
var err error = nil

var mplusNormalFont *text.GoXFace

var lastText []string = make([]string, 0, 5)

const FontSize = 16
const YPad = FontSize * 3 / 2
const XPad = FontSize * 2 / 3

func ProcessUserLog(g *Game, screen *ebiten.Image) {
	if userLogImg == nil {
		userLogImg, _, err = ebitenutil.NewImageFromFile("assets/UIPanel.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	if mplusNormalFont == nil {
		tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
		if err != nil {
			log.Fatal(err)
		}

		const dpi = 72
		var fontface font.Face
		fontface, err = opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    FontSize,
			DPI:     dpi,
			Hinting: font.HintingFull,
		})
		if err != nil {
			log.Fatal(err)
		}
		mplusNormalFont = text.NewGoXFace(fontface)
	}

	gd := NewGameData()

	uiLocation := (gd.ScreenHeight - gd.UIHeight) * gd.TileHeight
	var fontX = XPad
	var fontY = uiLocation + YPad

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(0.), float64(uiLocation))
	screen.DrawImage(userLogImg, op)
	tmpMessages := make([]string, 0, 5)
	anyMessages := false

	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.AttackMessage != "" {
			tmpMessages = append(tmpMessages, messages.AttackMessage)
			anyMessages = true
			//fmt.Printf(messages.AttackMessage)
			messages.AttackMessage = ""
		}
	}

	for _, m := range g.World.Query(g.WorldTags["messengers"]) {
		messages := m.Components[userMessage].(*UserMessage)
		if messages.DeadMessage != "" {
			tmpMessages = append(tmpMessages, messages.DeadMessage)
			anyMessages = true

			messages.DeadMessage = ""
			g.World.DisposeEntity(m.Entity)
		}

		if messages.GameStateMessage != "" {
			tmpMessages = append(tmpMessages, messages.GameStateMessage)
			anyMessages = true
		}
	}

	if anyMessages {
		lastText = tmpMessages
	}

	for _, msg := range lastText {
		if msg != "" {
			//text.Draw(screen, msg, mplusNormalFont, fontX, fontY, color.White)
			opts := &text.DrawOptions{}
			opts.GeoM.Translate(float64(fontX), float64(fontY))
			opts.ColorScale.ScaleWithColor(color.White)

			text.Draw(screen, msg, mplusNormalFont, opts)
			fontY += YPad
		}
	}
}

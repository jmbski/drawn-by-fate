package main

import (
	"image/color"
	"log"

	"github.com/RAshkettle/rrogue/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/mlange-42/ark/ecs"
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
		userLogImg = GetImage("UIPanel.png")
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
	anyMessages := false
	toRemove := []ecs.Entity{}
	atkMsgs := []string{}
	dedMsgs := []string{}
	gstMsgs := []string{}

	query := g.Context.Messengers.Query()

	for query.Next() {
		messages := query.Get()

		if messages.AttackMessage != "" {
			atkMsgs = append(atkMsgs, messages.AttackMessage)
			anyMessages = true
			messages.AttackMessage = ""
		}

		if messages.DeadMessage != "" {
			dedMsgs = append(dedMsgs, messages.DeadMessage)
			anyMessages = true

			messages.DeadMessage = ""

			//g.World.RemoveEntity(query.Entity())
			toRemove = append(toRemove, query.Entity())
		}

		if messages.GameStateMessage != "" {
			gstMsgs = append(gstMsgs, messages.GameStateMessage)
			anyMessages = true
		}
	}

	if anyMessages {
		lastText = append(lastText, atkMsgs...)
		lastText = append(lastText, dedMsgs...)
		lastText = append(lastText, gstMsgs...)
		for len(lastText) > 5 {
			lastText = lastText[1:]
		}
	}

	if len(toRemove) > 0 {
		for _, entity := range toRemove {
			g.World.RemoveEntity(entity)
		}
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

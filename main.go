package main

import (
	"drawn-by-fate/internal/engine"
	"drawn-by-fate/pkg/config"
	"drawn-by-fate/pkg/models"
	"drawn-by-fate/pkg/utils"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	arkserde "github.com/mlange-42/ark-serde"
)

var cfg = config.CurrentSettings

func main() {
	game := engine.NewGame()
	ebiten.SetWindowSize(config.GetScreenSize())
	ebiten.SetWindowTitle("Drawn By Fate - Vertical Slice")

	tmp, err := models.LoadEntityTemplates("core")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tmp.Mobs[0])
	if jsonData, err := arkserde.Serialize(game.World); err != nil {
		fmt.Println(err)
	} else {
		if err := utils.WriteBytesToFile(jsonData, "./world.json"); err != nil {
			fmt.Println(err)
		}
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

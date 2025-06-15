package imgs

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func GetImage(name string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile("assets/" + name)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

package systems

import (
	"drawn-by-fate/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func ProcessRenderables(game *engine.Game, screen *ebiten.Image) {
	query := game.Context.Renderables.Query()

	for query.Next() {
		render := query.Get()
		// TODO: add some checks

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(render.X, render.Y)
		screen.DrawImage(render.Image, op)
	}
}

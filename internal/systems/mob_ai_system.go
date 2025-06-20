package systems

import (
	cmp "drawn-by-fate/internal/components"
	"drawn-by-fate/pkg/models"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

type MobAiSystem struct {
	filter      *ecs.Filter4[cmp.Renderable, cmp.Velocity, cmp.BaseMovementSpeed, cmp.AiControlled]
	playerQuery *ecs.Filter2[cmp.InputControlled, cmp.Renderable]
}

func NewMobAiSystem(world *ecs.World) *MobAiSystem {
	return &MobAiSystem{
		filter:      ecs.NewFilter4[cmp.Renderable, cmp.Velocity, cmp.BaseMovementSpeed, cmp.AiControlled](world).Register(),
		playerQuery: ecs.NewFilter2[cmp.InputControlled, cmp.Renderable](world).Register(),
	}
}

func (s *MobAiSystem) Update() {

	playerPositions := make([]models.Vec2, 0)
	pQuery := s.playerQuery.Query()
	for pQuery.Next() {
		_, render := pQuery.Get()
		playerPositions = append(playerPositions, models.NewVec2(render.X, render.Y))
	}

	query := s.filter.Query()

	for query.Next() {
		render, vel, speed, ai := query.Get()
		pos := models.NewVec2(render.X, render.Y)
		dir := models.NewVec2(vel.X, vel.Y)

		switch ai.MovePattern {
		case cmp.AiMvToNearest:
			if len(playerPositions) == 0 {
				continue
			}

			var target models.Vec2
			minDist := float64(math.MaxFloat64)
			for _, p := range playerPositions {
				dx := render.X - p.X
				dy := render.Y - p.Y

				distSq := dx*dx + dy*dy
				if distSq < minDist {
					minDist = distSq
					target = p
				}
			}

			dir = target.Diff(pos)
			mag := dir.Len()
			if mag > 0 {
				vel.X = dir.X / mag * speed.Value
				vel.Y = dir.Y / mag * speed.Value
			} else {
				vel.X, vel.Y = 0, 0
			}
		case cmp.AiMvRandom:
			ai.TickCountDown--

			if ai.TickCountDown <= 0 {
				angle := rand.Float64() * 2 * math.Pi
				ai.DirX = math.Cos(angle)
				ai.DirY = math.Sin(angle)
				ai.TickCountDown = rand.Intn(30) + 20
			}
			dir.X = ai.DirX * speed.Value
			dir.Y = ai.DirY * speed.Value
		}

		render.X += dir.X / float64(ebiten.TPS())
		render.Y += dir.Y / float64(ebiten.TPS())
	}
}

func SmoothVelocity(dx, dy, speed float64) (float64, float64) {
	mag := math.Hypot(dx, dy)
	if mag != 0 {
		dx = (dx / mag) * speed
		dy = (dy / mag) * speed
	}
	return dx, dy
}

func GetRandVelocity(speed float64) (x, y float64) {
	xRand, yRand := rand.Float64(), rand.Float64()
	dx, dy := 0.0, 0.0

	if xRand > 0.5 {
		dx = 1.0
	} else {
		dx = -1.0
	}

	if yRand > 0.5 {
		dy = 1.0
	} else {
		dy = -1.0
	}

	return SmoothVelocity(dx, dy, speed)
}

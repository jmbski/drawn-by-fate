package engine

import (
	"drawn-by-fate/internal/core"
	"drawn-by-fate/internal/loaders"
	"drawn-by-fate/internal/scenes"
	"drawn-by-fate/internal/systems"
	"drawn-by-fate/pkg/config"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/mlange-42/ark/ecs"
)

var currentScene = scenes.CurrentScene()
var cfg = config.CurrentSettings
var coreInputName = "localPlayerInput"

type Game struct {
	World            *ecs.World
	Context          *core.WorldContext
	Config           config.GameConfig
	CorePaths        config.CommonPaths
	InputSystems     systems.SystemMap
	RenderSystem     *systems.RenderSystem
	PlayerMoveSystem *systems.PlayerMovementSystem
	AiMoveSystem     *systems.MobAiSystem
	AssetLoader      *loaders.AssetLoader
}

func (g *Game) Update() error {
	g.InputSystems.UpdateAll()
	g.PlayerMoveSystem.Update()
	g.AiMoveSystem.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.RenderSystem.Draw(screen, g.World)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nMoveSpeed: %0.2f\n", ebiten.ActualTPS(), ebiten.ActualFPS(), g.PlayerMoveSystem.LatestSpeed))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.GetScreenSize()
}

func NewGame() *Game {
	// TODO: set up and import NewWorld()/InitWorld() function in the core package
	world := core.InitializeWorld()
	is := systems.NewInputSystem()

	return &Game{
		World:            world,
		Context:          core.NewWorldContext(world),
		Config:           cfg,
		CorePaths:        config.CorePaths,
		InputSystems:     systems.SystemMap{coreInputName: is},
		RenderSystem:     systems.NewRenderSystem(world),
		PlayerMoveSystem: systems.NewPlayerMovementSystem(is, world),
		AiMoveSystem:     systems.NewMobAiSystem(world),
		AssetLoader:      loaders.NewAssetLoader(),
	}
}

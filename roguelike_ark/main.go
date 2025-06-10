package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/ark/ecs"
)

type ActorInstance struct {
	IsPlayer  bool
	Image     *ebiten.Image
	IsMovable bool
	Pos       *Position
	Health    *Health
	Weapon    *MeleeWeapon
	Armor     *Armor
	Name      *Name
	Message   *UserMessage
}

// Game holds all data the entire game will need.
type Game struct {
	Map         GameMap
	World       *ecs.World
	Context     *WorldContext
	Turn        TurnState
	TurnCounter int
}

// NewGame creates a new Game Object and initializes the data
func NewGame() *Game {
	g := &Game{}
	g.Map = NewGameMap()
	world, context := InitializeWorld(g.Map.CurrentLevel)
	g.Context = context
	g.World = world
	g.Turn = PlayerTurn
	g.TurnCounter = 0

	return g
}

func (g *Game) PlayerInstance(entity ecs.Entity) *ActorInstance {
	mapper := g.GetPlayers()
	if !mapper.HasAll(entity) {
		// todo: add error handling
		return nil
	}
	// Player, Renderable, Movable, Position, Health, MeleeWeapon, Armor, Name, UserMessage
	_, r, m, p, h, w, a, n, msg := mapper.Get(entity)
	return &ActorInstance{
		IsPlayer:  true,
		Image:     r.Image,
		IsMovable: m != nil,
		Pos:       p,
		Health:    h,
		Weapon:    w,
		Armor:     a,
		Name:      n,
		Message:   msg,
	}
}

func (g *Game) MonsterInstance(entity ecs.Entity) *ActorInstance {
	mapper := g.GetMobs()
	if !mapper.HasAll(entity) {
		return nil
	}
	// Monster, Renderable, Position, Health, MeleeWeapon, Armor, Name, UserMessage
	_, r, p, h, w, a, n, msg := mapper.Get(entity)
	return &ActorInstance{
		IsPlayer:  false,
		Image:     r.Image,
		IsMovable: true,
		Pos:       p,
		Health:    h,
		Weapon:    w,
		Armor:     a,
		Name:      n,
		Message:   msg,
	}
}

func (g *Game) CombatantInstance(entity ecs.Entity) *ActorInstance {

	if g.GetPlayers().HasAll(entity) {
		return g.PlayerInstance(entity)
	}

	if g.GetMobs().HasAll(entity) {
		return g.MonsterInstance(entity)
	}

	return nil
}

func (g *Game) GetPlayers() *PlayerMap {
	return GetPlayers(g.World)
}

func (g *Game) GetMobs() *MobMap {
	return GetMobs(g.World)
}

// Update is called each tic.
func (g *Game) Update() error {
	g.TurnCounter++
	if g.Turn == PlayerTurn && g.TurnCounter > 15 {
		TakePlayerAction(g)
	}

	if g.Turn == MonsterTurn {
		UpdateMonster(g)
	}
	return nil
}

// Draw is called each draw cycle and is where we will blit.
func (g *Game) Draw(screen *ebiten.Image) {
	//Draw the Map
	level := g.Map.CurrentLevel
	level.DrawLevel(screen)
	ProcessRenderables(g, level, screen)
	ProcessUserLog(g, screen)
	ProcessHUD(g, screen)
}

// Layout will return the screen dimensions.
func (g *Game) Layout(w, h int) (int, int) {
	gd := NewGameData()
	return gd.TileWidth * gd.ScreenWidth, gd.TileHeight * gd.ScreenHeight
}

func main() {
	g := NewGame()

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Tower")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

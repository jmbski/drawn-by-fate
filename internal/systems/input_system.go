package systems

import (
	"fmt"

	qi "github.com/quasilyte/ebitengine-input" // Renamed for clarity
)

// Action represents abstract game actions.
type Action = qi.Action

const (
	ActionMoveUp Action = iota
	ActionMoveDown
	ActionMoveLeft
	ActionMoveRight
)

// InputSystem manages all input handlers.
type InputSystem struct {
	System      qi.System
	PlayerInput *qi.Handler
}

func (s InputSystem) Update() {
	s.System.Update()
}

type TestIS struct {
	InputSystem qi.System
}

// NewInputSystem creates and initializes the input system.
func NewInputSystem() *InputSystem {
	is := &InputSystem{}
	fmt.Println(qi.SystemConfig{DevicesEnabled: qi.AnyDevice})
	is.System.Init(qi.SystemConfig{
		DevicesEnabled: qi.AnyDevice,
	})

	// Define keymap for player actions
	playerKeymap := qi.Keymap{
		ActionMoveUp:    {qi.KeyW, qi.KeyUp, qi.KeyGamepadUp},
		ActionMoveDown:  {qi.KeyS, qi.KeyDown, qi.KeyGamepadDown},
		ActionMoveLeft:  {qi.KeyA, qi.KeyLeft, qi.KeyGamepadLeft},
		ActionMoveRight: {qi.KeyD, qi.KeyRight, qi.KeyGamepadRight},
	}
	is.PlayerInput = is.System.NewHandler(0, playerKeymap) // Device ID 0 for primary player

	return is
}

type ISMap map[string]*InputSystem

func (s ISMap) UpdateAll() {
	for _, is := range s {
		is.Update()
	}
}

func NewISMap(names ...string) SystemMap {
	isMap := make(SystemMap, len(names))
	for _, name := range names {
		isMap[name] = NewInputSystem()
	}

	return isMap
}

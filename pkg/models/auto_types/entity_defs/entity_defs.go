// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    entityDefs, err := UnmarshalEntityDefs(bytes)
//    bytes, err = entityDefs.Marshal()

package entity_defs

import "encoding/json"

func UnmarshalEntityDefs(data []byte) (EntityDefs, error) {
	var r EntityDefs
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *EntityDefs) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type EntityDefs struct {
	Mobs    []MobEntity    `json:"mobs,omitempty" yaml:"mobs,omitempty"`
	Players []PlayerEntity `json:"players,omitempty" yaml:"players,omitempty"`
}

type MobEntity struct {
	ActorID           *string           `json:"actor_id,omitempty" yaml:"actor_id,omitempty"`
	ActorType         MobActorType      `json:"actor_type" yaml:"actor_type"`
	AIControlled      AIControlled      `json:"ai_controlled" yaml:"ai_controlled"`
	BaseMovementSpeed BaseMovementSpeed `json:"base_movement_speed" yaml:"base_movement_speed"`
	MetaData          MetaData          `json:"meta_data" yaml:"meta_data"`
	Renderable        Renderable        `json:"renderable" yaml:"renderable"`
	Velocity          Vector2           `json:"velocity" yaml:"velocity"`
}

type AIControlled struct {
	DirX          *float64    `json:"dir_x,omitempty" yaml:"dir_x,omitempty"`
	DirY          *float64    `json:"dir_y,omitempty" yaml:"dir_y,omitempty"`
	MovePattern   MovePattern `json:"move_pattern" yaml:"move_pattern"`
	TickCountdown *float64    `json:"tick_countdown,omitempty" yaml:"tick_countdown,omitempty"`
}

type BaseMovementSpeed struct {
	Value float64 `json:"value" yaml:"value"`
}

type MetaData struct {
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	UUID *string `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type Renderable struct {
	Image string  `json:"image" yaml:"image"`
	X     float64 `json:"x" yaml:"x"`
	Y     float64 `json:"y" yaml:"y"`
}

type Vector2 struct {
	X *float64 `json:"x,omitempty" yaml:"x,omitempty"`
	Y *float64 `json:"y,omitempty" yaml:"y,omitempty"`
}

type PlayerEntity struct {
	ActorID           string            `json:"actor_id" yaml:"actor_id"`
	ActorType         PlayerActorType   `json:"actor_type" yaml:"actor_type"`
	BaseMovementSpeed BaseMovementSpeed `json:"base_movement_speed" yaml:"base_movement_speed"`
	InputControlled   bool              `json:"input_controlled" yaml:"input_controlled"`
	MetaData          MetaData          `json:"meta_data" yaml:"meta_data"`
	Renderable        Renderable        `json:"renderable" yaml:"renderable"`
	Velocity          Vector2           `json:"velocity" yaml:"velocity"`
}

type MovePattern string

const (
	AIRandom    MovePattern = "ai_random"
	AIToNearest MovePattern = "ai_to_nearest"
)

type MobActorType string

const (
	Monster MobActorType = "monster"
)

type PlayerActorType string

const (
	Player PlayerActorType = "player"
)

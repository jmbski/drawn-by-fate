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
	Mobs    []MobEntity    `json:"mobs,omitempty"`
	Players []PlayerEntity `json:"players,omitempty"`
}

type MobEntity struct {
	ActorID           *string           `json:"actor_id,omitempty"`
	ActorType         MobActorType      `json:"actor_type"`
	AIControlled      AIControlled      `json:"ai_controlled"`
	BaseMovementSpeed BaseMovementSpeed `json:"base_movement_speed"`
	MetaData          MetaData          `json:"meta_data"`
	Renderable        Renderable        `json:"renderable"`
	Velocity          Vector2           `json:"velocity"`
}

type AIControlled struct {
	DirX          *float64    `json:"dir_x,omitempty"`
	DirY          *float64    `json:"dir_y,omitempty"`
	MovePattern   MovePattern `json:"move_pattern"`
	TickCountdown *float64    `json:"tick_countdown,omitempty"`
}

type BaseMovementSpeed struct {
	Value float64 `json:"value"`
}

type MetaData struct {
	Name *string `json:"name,omitempty"`
	UUID *string `json:"uuid,omitempty"`
}

type Renderable struct {
	Image string  `json:"image"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

type Vector2 struct {
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
}

type PlayerEntity struct {
	ActorID           string            `json:"actor_id"`
	ActorType         PlayerActorType   `json:"actor_type"`
	BaseMovementSpeed BaseMovementSpeed `json:"base_movement_speed"`
	InputControlled   bool              `json:"input_controlled"`
	MetaData          MetaData          `json:"meta_data"`
	Renderable        Renderable        `json:"renderable"`
	Velocity          Vector2           `json:"velocity"`
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

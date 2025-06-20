// Code generated from JSON Schema using quicktype. DO NOT EDIT.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    gameConfig, err := UnmarshalGameConfig(bytes)
//    bytes, err = gameConfig.Marshal()

package game_config

import "encoding/json"

func UnmarshalGameConfig(data []byte) (GameConfig, error) {
	var r GameConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *GameConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Main config file for general game options
type GameConfig struct {
	// List of mod ids that have been manually enabled by the user to be loaded into the current                
	// game context                                                                                             
	EnabledModIDS                                                                               []string        `json:"enabled_mod_ids" yaml:"enabled_mod_ids"`
	// Option to set how the game is rendered on screen (fullscreen, borderless window, or                      
	// window)                                                                                                  
	ScreenMode                                                                                  ScreenMode      `json:"screen_mode" yaml:"screen_mode"`
	// Height/width for the overall screen                                                                      
	ScreenResolution                                                                            HeightWidthPair `json:"screen_resolution" yaml:"screen_resolution"`
}

// Height/width for the overall screen
//
// Width and Height pairs for different visual elements such as screen size, screen
// resolution, etc..
type HeightWidthPair struct {
	Height int64 `json:"height" yaml:"height"`
	Width  int64 `json:"width" yaml:"width"`
}

// Option to set how the game is rendered on screen (fullscreen, borderless window, or
// window)
type ScreenMode string

const (
	BorderlessWindow ScreenMode = "borderlessWindow"
	Fullscreen       ScreenMode = "fullscreen"
	Window           ScreenMode = "window"
)

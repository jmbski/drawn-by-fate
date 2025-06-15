package config

type GameConfig struct {
	ScreenWidth  int // Width in Pixels
	ScreenHeight int // Height in pixels
}

func NewGameConfig() GameConfig {
	return GameConfig{
		ScreenWidth:  640,
		ScreenHeight: 480,
	}
}

var CurrentSettings = GameConfig{
	ScreenWidth:  640,
	ScreenHeight: 480,
}

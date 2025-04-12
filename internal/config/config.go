package config

type GameConfig struct {
	ScreenWidth  int
	ScreenHeight int
}

var CurrentSettings = GameConfig{
	ScreenWidth:  640,
	ScreenHeight: 480,
}

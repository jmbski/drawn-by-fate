package scenes

var currentScene GameScene

func CurrentScene() GameScene {
	return currentScene
}

func SwitchScene(s GameScene) {
	currentScene = s
}

package loaders

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ErrAssetNotFound = errors.New("asset not found")
)

// GetImageAsset loads an image asset given a POSIX-style relative path.
func (l *AssetLoader) GetImageAsset(assetPath string) (*ebiten.Image, error) {
	fullPath := filepath.Join(l.Paths.Images, assetPath)
	fmt.Println("asset path", fullPath)
	// Check if the file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		return nil, ErrAssetNotFound
	}

	// Load image using ebitenutil
	img, _, err := ebitenutil.NewImageFromFile(fullPath)
	if err != nil {
		return nil, err
	}

	return img, nil
}

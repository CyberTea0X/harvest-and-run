package assets

import (
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"path"
)

const (
	DefaultImagesPath = "./assets/images"
	DefaultFontsPath  = "./assets/fonts"
)

type Assets struct {
	imageAssets map[string]*ebiten.Image
	fontAssets  map[string]*truetype.Font
	imagesPath  string
	fontsPath   string
}

func New() *Assets {
	a := new(Assets)
	a.imageAssets = make(map[string]*ebiten.Image)
	a.fontAssets = make(map[string]*truetype.Font)
	a.imagesPath = DefaultImagesPath
	a.fontsPath = DefaultFontsPath
	return a
}

// Loads image assets or retrieves loaded asset
//
// Prepends imagesPath to path variable,
// so you don't have to specify whole path to images each time
func (a *Assets) Image(imagePath string) (*ebiten.Image, error) {
	wholePath := path.Join(a.imagesPath, imagePath)
	image, e := a.imageAssets[wholePath]
	if e {
		return image, nil
	}
	image, _, err := ebitenutil.NewImageFromFile(wholePath)
	if err != nil {
		return nil, err
	}
	a.imageAssets[wholePath] = image
	return image, nil
}

// Loads font assets or retrieves loaded asset
//
// Prepends fontsPath to path variable,
// so you don't have to specify whole path to fonts each time
func (a *Assets) Font(fontPath string) (*truetype.Font, error) {
	wholePath := path.Join(a.fontsPath, fontPath)
	font, e := a.fontAssets[wholePath]
	if e {
		return font, nil
	}
	rawFont, err := os.ReadFile(wholePath)
	if err != nil {
		return nil, err
	}
	font, err = truetype.Parse(rawFont)
	if err != nil {
		return nil, err
	}
	return font, nil
}

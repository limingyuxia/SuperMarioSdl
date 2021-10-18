package gameMap

import (
	"log"
	"superMario/model"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// 图片转换为纹理
func ImgToTexture(imagePath string, renderer *sdl.Renderer) (*model.ImgTextureInfo, error) {
	// 加载图片
	gameImg, err := img.Load(imagePath)
	if err != nil {
		log.Print("load game image failed: ", err)
		return nil, err
	}
	defer gameImg.Free()

	// 去除杂色
	gameImg.SetColorKey(true, sdl.MapRGB(gameImg.Format, 255, 0, 255))

	// 创建纹理
	textrue, err := renderer.CreateTextureFromSurface(gameImg)
	if err != nil {
		log.Print("create texture failed: ", err)
		return nil, err
	}

	return &model.ImgTextureInfo{
		ImgTexture: textrue,
		Width:      gameImg.W,
		Height:     gameImg.H,
	}, nil
}

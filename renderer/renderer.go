package renderer

import (
	"superMario/model"

	"github.com/veandco/go-sdl2/sdl"
)

func RendererMap(renderer *sdl.Renderer, role *model.ImgTextureInfo, barriers []*model.ImgTextureInfo) {

	// 渲染障碍物
	for _, barrier := range barriers {
		renderer.Copy(barrier.ImgTexture, nil, &sdl.Rect{
			X: barrier.PosX,
			Y: barrier.PosY,
			W: barrier.Width,
			H: barrier.Height,
		})
	}

	// 渲染人物
	renderer.Copy(role.ImgTexture, nil, &sdl.Rect{
		X: role.PosX,
		Y: role.PosY,
		W: role.Width,
		H: role.Height,
	})
}

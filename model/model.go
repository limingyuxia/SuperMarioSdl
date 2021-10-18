package model

import "github.com/veandco/go-sdl2/sdl"

const WindowHeight int32 = 448
const WindowWidth int32 = 800

type ImgTextureInfo struct {
	ImgTexture *sdl.Texture `json:"img_texture"` // 图片转换后的纹理
	Width      int32        `json:"width"`       // 图片宽度
	Height     int32        `json:"height"`      // 图片高度
	PosX       int32        `json:"pos_x"`       // 在地图中的x轴坐标
	PosY       int32        `json:"pos_y"`       // 在地图中的y轴坐标
}

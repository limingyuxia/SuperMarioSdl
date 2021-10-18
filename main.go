package main

import (
	"log"
	"superMario/collision"
	"superMario/gameMap"
	"superMario/model"
	rendererMap "superMario/renderer"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var window *sdl.Window

func init() {
	// sdl 初始化
	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_TIMER | sdl.INIT_AUDIO); err != nil {
		log.Print("sdl init error: ", err)
		return
	}

	// 创建窗口
	var err error
	window, err = sdl.CreateWindow("超级玛丽", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		model.WindowWidth, model.WindowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Print("create window error: ", err)
		return
	}

	// 设置应用图标
	iconFile := "./resource/images/ico.bmp"
	iconBmp, err := sdl.LoadBMP(iconFile)
	if err != nil {
		log.Print("load icon failed: ", err)
		return
	}
	defer iconBmp.Free()

	iconBmp.SetColorKey(true, sdl.MapRGB(iconBmp.Format, 255, 0, 255))
	window.SetIcon(iconBmp)
}

func main() {
	defer sdl.Quit()
	defer window.Destroy()

	// 初始化音频
	mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 2048)

	// backMusic, err := mix.LoadWAV("./resource/sounds/overworld.wav")
	// if err != nil {
	// 	log.Print("load overworld music failed: ", err)
	// 	return
	// }
	// _, err = backMusic.Play(2, -1)
	// if err != nil {
	// 	log.Print("play overworld music failed: ", err)
	// 	return
	// }

	// 设置图形渲染硬件加速
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Print("create renderer failed: ", err)
		return
	}
	defer renderer.Destroy()

	var moveDistince int32 = 12

	brickredPath := "./resource/images/brickred.bmp"
	brickred, err := gameMap.ImgToTexture(brickredPath, renderer)
	if err != nil {
		return
	}
	brickred.PosX = 256
	brickred.PosY = 96

	marioPath := "./resource/images/mario/mario2_move1.bmp"
	mario, err := gameMap.ImgToTexture(marioPath, renderer)
	if err != nil {
		return
	}
	mario.PosX = 32
	mario.PosY = 128

	var barriers []*model.ImgTextureInfo
	barriers = append(barriers, brickred)

	running := true
	for running {
		sdl.Delay(15)
		renderer.Clear()

		rendererMap.RendererMap(renderer, mario, barriers)

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				switch t.Keysym.Scancode {
				case sdl.SCANCODE_D:
					mario.PosX += moveDistince
				case sdl.SCANCODE_A:
					mario.PosX -= moveDistince
				case sdl.SCANCODE_S:
					mario.PosY += moveDistince
				case sdl.SCANCODE_W:
					mario.PosY -= moveDistince
				}

				collision.CheckCollison(mario, barriers)
			}
		}

		renderer.Present()
	}
}

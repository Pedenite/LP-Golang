package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl" //go get ...
)

const winWidth int = 800
const winHeight int = 600

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

type ball struct {
	pos    //isto eh tipo uma heranca em go, em que todos os atributos do struct serao herdados por este
	radius int
	yv     float32
	xv     float32
	color  color //o go permite se ter uma variavel com o mesmo nome q o tipo
}

type paddle struct {
	pos
	w     int
	h     int
	color color
}

func setPixel(x, y int, c color, pixel []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixel)-4 && index >= 0 {
		pixel[index] = c.r
		pixel[index+1] = c.g
		pixel[index+2] = c.b
	}
}

func main() {

	///////////////////////////////////preparando a janela a ser aberta e poder printar pixels/////////////////////////////
	window, err := sdl.CreateWindow("Testando SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy() //defer faz so ser executado ao fim da funcao

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight)) //ARGB -> A remete a transparencia
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()
	////////////////////////////////////////////fim preparacao///////////////////////////////////////////

	pixels := make([]byte, winWidth*winHeight*4)

	//sdl.Delay(2000)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			//setPixel(x, y, color{255, 0, 0}, pixels) //torna os pixels vermelhos
			setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
		}
	}

	tex.Update(nil, pixels, winWidth*4) //esse 4 significa quantos bytes por pixel -> 1 R, 1 G, 1 B e 1 A
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	sdl.Delay(2000)
}

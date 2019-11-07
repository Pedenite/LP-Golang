package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth = 800
const winHeight = 600
const velPlayer = 100
const velTiro = 500

var paredes = 100

type cor struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

type wall struct {
	pos
	cor
	l int
}

func (wall *wall) draw(pixels []byte) {
	inicioX := int(wall.x) - wall.l
	inicioY := 0

	for y := 0; y < winHeight; y++ {
		for x := 0; x < wall.l; x++ {
			setPixel(inicioX+x, inicioY+y, wall.cor, pixels)
		}
	}
}

type aviao struct {
	pos
	cor
	a int
	l int
	v float32
}

func (aviao *aviao) draw(pixels []byte) {
	inicioX := int(aviao.x) - aviao.l/2
	inicioY := int(aviao.y) - aviao.a/2
	for y := 0; y < aviao.a; y++ {
		for x := 0; x < aviao.l; x++ {
			setPixel(inicioX+x, inicioY+y, aviao.cor, pixels)
		}
	}
}

func (aviao *aviao) update(keyState []uint8, wall1 *wall, wall2 *wall, tiro *tiro, pixels []uint8, elapsedTime float32) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		aviao.y -= aviao.v * elapsedTime
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		aviao.y += aviao.v * elapsedTime
	}
	if keyState[sdl.SCANCODE_LEFT] != 0 {
		aviao.x -= aviao.v * elapsedTime
	}
	if keyState[sdl.SCANCODE_RIGHT] != 0 {
		aviao.x += aviao.v * elapsedTime
	}

	if aviao.x-float32(aviao.l/2) < wall1.x || aviao.x+float32(aviao.l/2) > wall2.x-float32(wall2.l) {
		aviao.x = 400
		aviao.y = 450
	}

	if keyState[sdl.SCANCODE_SPACE] != 0 {
		tiro.x = aviao.x
		tiro.y = aviao.y - float32(aviao.a)/2
	}
}

type tiro struct {
	pos
	cor
	a int
	l int
	v float32
}

func (tiro *tiro) draw(pixels []uint8) {
	inicioX := int(tiro.x) - tiro.l
	inicioY := int(tiro.y) - tiro.a
	for y := 0; y < tiro.a; y++ {
		for x := 0; x < tiro.l; x++ {
			setPixel(inicioX+x, inicioY+y, tiro.cor, pixels)
		}
	}
}
func (tiro *tiro) update(elapsedTime float32) {
	tiro.y -= tiro.v * elapsedTime
	if tiro.y < 0 {
		tiro.x = 0
		tiro.y = 0
	}
}

func clear(pixels []byte, c cor) {
	for i := range pixels {
		switch i % 4 {
		case 0:
			pixels[i] = c.r
		case 1:
			pixels[i] = c.g
		case 2:
			pixels[i] = c.b
		default:
			pixels[i] = 0
		}
	}
}

func setPixel(x, y int, c cor, pixel []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixel)-4 && index >= 0 {
		pixel[index] = c.r
		pixel[index+1] = c.g
		pixel[index+2] = c.b
	}
}
func main() {
	window, err := sdl.CreateWindow("River Raid", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
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

	pixels := make([]byte, winWidth*winHeight*4) //cria um slice para os pixels

	keyState := sdl.GetKeyboardState()

	var frameStart time.Time
	var elapsedTime float32 //time.Duration

	player := aviao{pos{winWidth / 2, winHeight * 3 / 4}, cor{150, 150, 0}, 20, 20, velPlayer}
	wall1 := wall{pos{100, 0}, cor{0, 100, 0}, paredes}
	wall2 := wall{pos{winWidth, 0}, cor{0, 100, 0}, paredes}
	tiro := tiro{pos{0, 0}, cor{255, 255, 255}, 3, 3, velTiro}

	for { //Game loop
		frameStart = time.Now()

		//necessario para input do teclado
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		clear(pixels, cor{0, 50, 100})
		wall1.draw(pixels)
		wall2.draw(pixels)
		player.draw(pixels)
		tiro.draw(pixels)

		player.update(keyState, &wall1, &wall2, &tiro, pixels, elapsedTime)
		tiro.update(elapsedTime)

		tex.Update(nil, pixels, winWidth*4) //esse 4 significa quantos bytes por pixel -> 1 R, 1 G, 1 B e 1 A
		renderer.Copy(tex, nil, nil)
		renderer.Present()

		elapsedTime = float32(time.Since(frameStart).Seconds())

		if elapsedTime < .005 { //max fps = 200
			sdl.Delay(5 - uint32(elapsedTime)/1000)
			elapsedTime = float32(time.Since(frameStart).Seconds())
		}
	}
}

package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl" //go get ...
)

const winWidth int = 800
const winHeight int = 600
const velPlayer float32 = 5
const initVelBall float32 = 5

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32
}

/////////////////////////////////////////Bola////////////////////////////////////////////

type ball struct {
	pos            //isto eh tipo uma heranca em go, em que todos os atributos do struct serao herdados por este
	radius int     //raio
	yv     float32 //velocidade coordenada y
	xv     float32 //velocidade coordenada x
	color  color   //o go permite se ter uma variavel com o mesmo nome q o tipo
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

func (ball *ball) update(paddle1 *paddle, paddle2 *paddle) {
	ball.x += ball.xv
	ball.y += ball.yv

	if int(ball.y)+ball.radius < 0 || int(ball.y)+ball.radius > winHeight { //collisao cima e baixo
		ball.yv = -ball.yv
	}
	if int(ball.x)+ball.radius < 0 || int(ball.x)+ball.radius > winWidth {
		ball.x = float32(winWidth) / 2
		ball.y = float32(winHeight) / 2
	}
	if ball.x-float32(ball.radius) < paddle1.x+float32(paddle1.w)/2 {
		if ball.y > paddle1.y-float32(paddle1.h)/2 && ball.y < paddle1.y+float32(paddle1.h)/2 {
			ball.xv = -ball.xv
		}
	}
	if ball.x+float32(ball.radius) > paddle2.x-float32(paddle2.w)/2 {
		if ball.y > paddle2.y-float32(paddle2.h)/2 && ball.y < paddle2.y+float32(paddle2.h)/2 {
			ball.xv = -ball.xv
		}
	}
}

/////////////////////////////////////fimBola/////////////////////////////////////////////////////

///////////////////////////////////Player////////////////////////////////////////////////////////
type paddle struct {
	pos         //posicao inicial
	w     int   //largura
	h     int   //altura
	color color //cor
}

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.w/2
	startY := int(paddle.y) - paddle.h/2 //pegara o pixel mais a esquerda da primeira linha de pixels

	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}
}

func (paddle *paddle) update1(keyState []uint8) {
	if keyState[sdl.SCANCODE_UP] != 0 {
		paddle.y -= velPlayer
	}
	if keyState[sdl.SCANCODE_DOWN] != 0 {
		paddle.y += velPlayer
	}
}
func (paddle *paddle) update2(keyState []uint8) {
	if keyState[sdl.SCANCODE_W] != 0 {
		paddle.y -= velPlayer
	}
	if keyState[sdl.SCANCODE_S] != 0 {
		paddle.y += velPlayer
	}
}

func (paddle *paddle) aiUpdate(ball *ball) {
	paddle.y = ball.y
}

////////////////////////////////////////////fimPlayer///////////////////////////////////////
func clear(pixels []byte) {
	for i := range pixels {
		pixels[i] = 0
	}
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
	window, err := sdl.CreateWindow("PONG", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
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
	////////////////////////////////////////////fim preparacao///////////////////////////////////////////

	keyState := sdl.GetKeyboardState()

	player1 := paddle{pos{50, 100}, 10, 75, color{255, 0, 255}}
	player2 := paddle{pos{float32(winWidth) - 50, 100}, 10, 75, color{255, 255, 0}}
	ball := ball{pos{float32(winWidth) / 2, float32(winHeight) / 2}, 10, initVelBall, initVelBall, color{255, 255, 255}}

	for { //Game loop
		//necessario para input do teclado
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		clear(pixels)

		player1.update1(keyState)
		player2.update2(keyState)
		ball.update(&player1, &player2)

		player1.draw(pixels)
		player2.draw(pixels)
		ball.draw(pixels)

		tex.Update(nil, pixels, winWidth*4) //esse 4 significa quantos bytes por pixel -> 1 R, 1 G, 1 B e 1 A
		renderer.Copy(tex, nil, nil)
		renderer.Present()
		sdl.Delay(16)
	}

	//sdl.Delay(2000)
}

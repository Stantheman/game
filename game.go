// author: Jacky Boen

package main

import (
	"fmt"
	"github.com/Stantheman/game/sound"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

var winTitle string = "Go-SDL2 Render"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var rect sdl.Rect

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	renderer.Clear()
	s, err := sound.Init()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sound: %v\n", err)
		os.Exit(3)
	}
	if err := s.Beep(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to beep: %v\n", err)
		os.Exit(4)
	}

	rect = sdl.Rect{0, 0, int32(winWidth), int32(winHeight)}
	for red := 0; red < 256; red++ {
		renderer.SetDrawColor(uint8(red), uint8(red), uint8(red), 255)
		renderer.FillRect(&rect)

		renderer.Present()

		sdl.Delay(15)

	}
	renderer.Destroy()
	window.Destroy()
}

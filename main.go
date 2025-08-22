package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Simple Clock")
	ebiten.SetVsyncEnabled(true)
	ebiten.SetTPS(30)
	ebiten.SetFullscreen(true)

	clock := NewClock()

	if err := ebiten.RunGame(clock); err != nil {
		log.Fatal(err)
	}
}

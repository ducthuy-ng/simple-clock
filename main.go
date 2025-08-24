package main

import (
	"fmt"
	"image"
	"time"

	"github.com/ducthuy-ng/simple-clock/programs"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
)

const (
	UPDATE_RATE_PER_SECOND = 60
	SCREEN_WIDTH           = 480
	SCREEN_HEIGHT          = 320
)

func main() {
	possiblePrograms := []programs.Program{
		programs.NewAnalogClockProgram(),
		programs.NewClockProgram(),
	}
	currentProgramIndex := 0

	driver.Main(func(s screen.Screen) {
		window, err := s.NewWindow(&screen.NewWindowOptions{
			Width:  SCREEN_WIDTH,
			Height: SCREEN_HEIGHT,
			Title:  "Simple Clock",
		})
		if err != nil {
			panic(fmt.Sprintf("failed to get new window: %v", err))
		}
		defer window.Release()

		for _, program := range possiblePrograms {
			buffer, err := s.NewBuffer(image.Point{X: SCREEN_WIDTH, Y: SCREEN_HEIGHT})
			if err != nil {
				panic(fmt.Sprintf("failed to create new buffer: %v", err))
			}
			defer buffer.Release()
			program.Init(buffer)
		}

		fpsTicker := time.NewTicker(time.Second / UPDATE_RATE_PER_SECOND)

		go func() {
			for {
				<-fpsTicker.C
				window.Send(paint.Event{})
			}
		}()

		for {
			program := possiblePrograms[currentProgramIndex%len(possiblePrograms)]

			e := window.NextEvent()
			switch e := e.(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case paint.Event:
				buffer := program.Draw()
				window.Upload(image.Point{0, 0}, buffer, buffer.Bounds())
				window.Publish()
			case key.Event:
				if e.Code == key.CodeEscape {
					return
				}
			case mouse.Event:
				if e.Button == mouse.ButtonLeft && e.Direction == mouse.DirPress {
					currentProgramIndex++
				}
			}
		}
	})
}

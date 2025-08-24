package programs

import (
	"image/color"
	"time"

	"math"

	"github.com/fogleman/gg"
	"golang.org/x/exp/shiny/screen"
)

type AnalogClockProgram struct {
	buffer screen.Buffer
	radius float64

	ScalingFactor float64
}

func NewAnalogClockProgram() *AnalogClockProgram {
	return &AnalogClockProgram{ScalingFactor: 1, radius: 100}
}

func (program *AnalogClockProgram) Init(buffer screen.Buffer) {
	program.buffer = buffer
}

func (program *AnalogClockProgram) Draw() screen.Buffer {
	program.fillBackground(color.RGBA{255, 255, 255, 255})
	program.drawClock()
	return program.buffer
}

func (program *AnalogClockProgram) Release() {
	panic("TODO: Implement")
}

func (program *AnalogClockProgram) fillBackground(color1 color.Color) {
	for x := 0; x < program.buffer.Size().X; x++ {
		for y := 0; y < program.buffer.Size().Y; y++ {
			program.buffer.RGBA().Set(x, y, color1)
		}
	}
}

func (program *AnalogClockProgram) drawClock() {
	currentTime := time.Now()

	second := float64(currentTime.Second())
	minute := float64(currentTime.Minute()) + second/60
	hour := float64(currentTime.Hour()%12) + minute/60

	secondPhi := second/60*2*math.Pi - (math.Pi / 2)
	minutePhi := minute/60*2*math.Pi - (math.Pi / 2)
	hourPhi := hour/12*2*math.Pi - (math.Pi / 2)

	actualRadius := program.radius * program.ScalingFactor
	screenSize := program.buffer.Size()
	ggContext := gg.NewContextForRGBA(program.buffer.RGBA())

	ggContext.SetColor(color.Black)
	ggContext.DrawCircle(float64(screenSize.X)/2, float64(screenSize.Y)/2, actualRadius)
	ggContext.Fill()

	ggContext.SetColor(color.White)
	ggContext.DrawCircle(float64(screenSize.X)/2, float64(screenSize.Y)/2, actualRadius-5)
	ggContext.Fill()

	// Draw second hand
	secondHandleLength := actualRadius - 15
	ggContext.SetColor(color.RGBA{255, 191, 0, 255})
	ggContext.SetLineWidth(3)
	ggContext.DrawLine(
		float64(screenSize.X)/2,
		float64(screenSize.Y)/2,
		float64(screenSize.X)/2+secondHandleLength*math.Cos(secondPhi),
		float64(screenSize.Y)/2+secondHandleLength*math.Sin(secondPhi),
	)
	ggContext.Stroke()

	// Draw minute hand
	minuteHandleLength := actualRadius - 20
	ggContext.SetColor(color.Black)
	ggContext.SetLineWidth(5)
	ggContext.DrawLine(
		float64(screenSize.X)/2,
		float64(screenSize.Y)/2,
		float64(screenSize.X)/2+minuteHandleLength*math.Cos(minutePhi),
		float64(screenSize.Y)/2+minuteHandleLength*math.Sin(minutePhi),
	)
	ggContext.Stroke()

	// Draw hour hand
	hourHandleLength := actualRadius - 30
	ggContext.SetColor(color.Black)
	ggContext.SetLineWidth(5)
	ggContext.DrawLine(
		float64(screenSize.X)/2,
		float64(screenSize.Y)/2,
		float64(screenSize.X)/2+hourHandleLength*math.Cos(hourPhi),
		float64(screenSize.Y)/2+hourHandleLength*math.Sin(hourPhi),
	)
	ggContext.Stroke()
}

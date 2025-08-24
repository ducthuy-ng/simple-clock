package programs

import "golang.org/x/exp/shiny/screen"

type Program interface {
	Init(screen.Buffer)
	Draw() screen.Buffer
	Release()
}

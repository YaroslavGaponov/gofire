package gofire

import (
	"math/rand"
)

const DEF_COLOR = 232

var PALLETE = [...]int{231, 230, 187, 185, 143, 142, 172, 166, 130, 124, 88, 1, 52, 234, 233, 232}

type Fire struct {
	term   *Term
	screen [][]int
}

func NewFire(width, height int) Fire {

	screen := make([][]int, height)
	for i := 0; i < height; i++ {
		screen[i] = make([]int, width)
	}

	return Fire{
		term:   &Term{Width: width, Height: height},
		screen: screen,
	}
}

func (fire *Fire) Draw() {

	for i := 1; i < fire.term.Height; i++ {
		for j := 1; j < fire.term.Width-1; j++ {
			n := rand.Intn(2)
			w := j - n + 1
			fire.screen[i][w] = fire.screen[i-1][j] + (n & 1)
		}
	}

	fire.term.Print(fire.term.Cursor(false))

	for i := 0; i < fire.term.Height; i++ {
		for j := 0; j < fire.term.Width; j++ {
			color := DEF_COLOR
			if fire.screen[i][j] < len(PALLETE) {
				color = PALLETE[fire.screen[i][j]]
			}
			fire.term.Print(fire.term.GotoXY(i, j) + fire.term.Backgroud(color))
		}
	}

	fire.term.Print(fire.term.Cursor(true))

	fire.term.Flush()
}

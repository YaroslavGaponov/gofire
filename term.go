package gofire

import (
	"fmt"
	"strings"
)

type Term struct {
	Width  int
	Height int
	buffer []string
}

func NewTerm(width, height int) *Term {
	term := Term{
		Width:  width,
		Height: height,
		buffer: []string{},
	}
	return &term
}

func (term *Term) Cursor(visible bool) string {
	if visible {
		return "\x1b[?25h"
	} else {
		return "\x1b[?25l"
	}
}

func (term *Term) GotoXY(x, y int) string {
	return fmt.Sprintf("\x1b[%d;%dH", term.Height-x, y)
}
func (term *Term) Backgroud(color int) string {
	return fmt.Sprintf("\x1b[48;5;%dm ", color)
}

func (term *Term) Print(s string) {
	term.buffer = append(term.buffer, s)
}

func (term *Term) Flush() {
	fmt.Print(strings.Join(term.buffer, ""))
	term.buffer = []string{}
}

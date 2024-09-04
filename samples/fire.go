package main

import (
	".."
	"time"
)

func main() {
	fire := gofire.NewFire(50, 30)
	for {
		fire.Draw()
		time.Sleep(100 * time.Millisecond)
	}
}

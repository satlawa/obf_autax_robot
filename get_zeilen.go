package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MoveMouse(150, 150)
	robotgo.Click()
	x, y := robotgo.GetMousePos()
	color := robotgo.GetPixelColor(x, y)
	color2 := robotgo.GetPxColor(x, y)

	zeilen := 0
	for i := 510; i < 700; i = i + 25 {
		//robotgo.MoveMouse(240, i)
		color = robotgo.GetPixelColor(265, i)
		color2 = robotgo.GetPxColor(265, i)
		//color := robotgo.GetPixelColor(222, i)
		fmt.Println("pos: 265 ", i)
		fmt.Println("color---- ", color)
		fmt.Println("color2---- ", color2)
		zeilen = zeilen + 1
	}
}

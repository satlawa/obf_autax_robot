package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

func main() {
	// gets the mouse coordinates
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	robotgo.Move(400, 400)

	// gets the mouse coordinates
	x, y = robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	robotgo.MoveMouse(100, 200)

	// gets the mouse coordinates
	x, y = robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	robotgo.ScrollMouse(10, "up")
	//robotgo.MouseClick("left", true) //double click
	robotgo.Click() //one click
	//robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	// gets the mouse coordinates
	x, y = robotgo.GetMousePos()
	fmt.Println("pos:", x, y)
}

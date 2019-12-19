package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type nutzung struct {
	X string
}

func new_nutz() {
	// move mouse to
	robotgo.MoveMouse(222, 506)
	//one left click
	robotgo.Click()
	// move mouse to
	robotgo.MoveMouse(222, 595)
	//one left click
	robotgo.Click()
}

func fill_nutz(array [9]nutzung) {
	// press "tab"
	for i := 0; i < len(array); i++ {
		robotgo.KeyTap("tab")
		robotgo.TypeStr(array[i].X)
	}
}

func new_bz() {
	// move mouse to
	robotgo.MoveMouse(822, 300)
	//one left click
	robotgo.Click()
}

func fill_bz(ba_art string, ba_anteil string) {

	robotgo.TypeStr(ba_art)
	robotgo.KeyTap("tab")
	robotgo.TypeStr(ba_anteil)
}

func main() {
	// gets the mouse coordinates
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	// insert nutz field
	nu1 := []nutzung{"DE", "1,4", "100", "20", "1", "1", "3", "4", "35"}
	new_nutz()
	fill_nutz(nu1[:])

	// insert nutz field
	new_bz()
	fill_bz("AH", "6")
	new_bz()
	fill_bz("EI", "4")

}

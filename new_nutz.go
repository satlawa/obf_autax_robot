package main

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

// next Waldort
func next_wo() {
	// move mouse to
	robotgo.MoveMouse(295, 100)
	//one left click
	robotgo.Click()
}

// make new Nutzung
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

// fill Nutzung with data
func fill_nutz(array []string) {
	// press "tab"
	for i := 0; i < len(array); i++ {
		robotgo.KeyTap("tab")
		robotgo.TypeStr(array[i])
	}
}

// make new Bestockungsziel
func new_bz() {
	// move mouse to
	robotgo.MoveMouse(822, 300)
	//one left click
	robotgo.Click()
}

// fill Bestockungsziel with data
func fill_bz(ba_art string, ba_anteil string) {

	robotgo.TypeStr(ba_art)
	robotgo.KeyTap("tab")
	robotgo.TypeStr(ba_anteil)
}

// not working
func get_zeilen() {
	zeilen := 0
	for i := 510; i < 600; i = i + 23 {
		robotgo.MoveMouse(222, i)
		color := robotgo.GetPixelColor(222, i)
		fmt.Println("color---- ", color)
		zeilen = zeilen + 1
	}

}

func main() {
	// gets the mouse coordinates
	x, y := robotgo.GetMousePos()
	fmt.Println("pos:", x, y)

	// insert nutz field
	nu1 := [9]string{"DE", "1,4", "100", "20", "1", "1", "3", "4", "35"}
	new_nutz()
	fill_nutz(nu1[:])

	// insert nutz field
	new_bz()
	fill_bz("AH", "6")
	new_bz()
	fill_bz("EI", "4")

	// next Teilfläche
	//next_wo()

	//new_bz()
	//fill_bz("EI", "6")
	//new_bz()
	//fill_bz("EI", "4")

	// insert nutz field
	//nu1 = [9]string{"DE", "1,4", "100", "20", "1", "1", "3", "4", "35"}
	//new_nutz()
	//fill_nutz(nu1[:])

}

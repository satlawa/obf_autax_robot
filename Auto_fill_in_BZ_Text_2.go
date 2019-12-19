package main

/*
program for automating TaxMaske
takes as input a csv file named data.csv
and fills data into Taxation Offline Client
version 0.2
*/

import (
	"fmt"
	"strings"
	//"context"
	"io/ioutil"
	"strconv"
	// external libreries
	dataframe "github.com/go-gota/gota/dataframe"
	"github.com/go-vgo/robotgo"
	//series "github.com/go-gota/gota/series"
)

func main() {

	// define map[type]type - define dictionary column : baumart
	//FI,TA,LA,KI,BU,EI,AH,ES,ER
	var map_ba = map[int]string{
		11: "FI",
		12: "TA",
		13: "LA",
		14: "KI",
		15: "BU",
		16: "EI",
		17: "AH",
		18: "ES",
		19: "ER",
	}

	// load data from csv file
	csvStr, err := ioutil.ReadFile("data_BZ_text.csv")
	check(err)

	// load string data 'dat' to dataframe
	df := dataframe.ReadCSV(strings.NewReader(string(csvStr)))
	df_row, _ := df.Dims()

	// print dataframe 'df' content
	fmt.Println(df)

	// loop rows
	for row := 0; row < df_row; row++ {

		// *** Text ***
		new_text()
		// loop columns 20 to 22 (Text)
		for col := 20; col < 23; col++ {
			text := df.Elem(row, col).String()
			if text != "0" {
				robotgo.TypeStr(text)
				if col < 22 {
					robotgo.KeyTap("enter")
				}
			}
		}

		// *** Maßnahme ***
		massnahme := df.Elem(row, 2).String()
		fmt.Println("Maßnahme" + massnahme)
		if massnahme != "0" {
			// get number of row in Schichtmerkmale
			ro := df.Elem(row, 1).String()
			rows_schicht, err := strconv.Atoi(ro)
			if err == nil {
				fmt.Println(rows_schicht)
			}

			new_nutz(rows_schicht)
			// loop columns 2 to 10 (Maßnahmenzeile)
			for col := 2; col < 11; col++ {
				// get elemet
				elem := df.Elem(row, col).String()
				// fill in data
				robotgo.KeyTap("tab")
				robotgo.TypeStr(elem)
			}
		}

		// *** Bestockungsziel ***
		fmt.Println("Bestockungsziel")
		// delete old
		delete_bz()
		// loop though columns 11 to 19 (BA-Bestockungsziele)
		for col := 11; col < 20; col++ {
			// get elemet
			elem := df.Elem(row, col).String()
			if elem != "0" {
				// get baumart from dictionary
				elem_ba := map_ba[col]
				// insert nutz field and fill in data
				new_bz()
				fill_bz(elem_ba, elem)
			}
		}
		// next Waldort
		next_wo()
	}
}

// Reading files requires checking most calls for errors. This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// next Waldort
func next_wo() {
	// move mouse to
	robotgo.MoveMouse(295, 90) // 295, 100
	//one left click
	robotgo.Click()
}

// make new Nutzung
func new_nutz(rows int) {
	coord := 560 + 22*(rows-1) // 575
	// move mouse to
	robotgo.MoveMouse(222, 495) // 222, 506
	//one left click
	robotgo.Click()
	// move mouse to
	robotgo.MoveMouse(222, coord)
	fmt.Println(coord)
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

// delete old Bestockungsziel
func delete_bz() {
	// move mouse to
	robotgo.MoveMouse(1470, 340) // 822, 300
	//one left click
	robotgo.Click()
	// move mouse to
	robotgo.MoveMouse(1500, 290) // 822, 300
	// 5 left clicks
	for i := 0; i < 5; i++ {
		robotgo.Click()
	}
}

// make new Bestockungsziel
func new_bz() {
	// move mouse to
	robotgo.MoveMouse(1470, 290) // 822, 300
	//one left click
	robotgo.Click()
}

// fill Bestockungsziel with data
func fill_bz(ba_art string, ba_anteil string) {

	robotgo.TypeStr(ba_art)
	robotgo.KeyTap("tab")
	robotgo.TypeStr(ba_anteil)
}

func new_text() {
	// move mouse to
	robotgo.MoveMouse(350, 350)
	//one left click
	robotgo.Click()
	robotgo.KeyTap("a", "ctrl")
	robotgo.KeyTap("delete")
}

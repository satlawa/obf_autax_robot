package main

/*
program for automating TaxMaske
takes as input a csv file named data.csv
and fills data into Taxation Offline Client
version 0.1
*/

import (
	"fmt"
	"strings"
	//"context"
	"io/ioutil"
	"strconv"

	dataframe "github.com/go-gota/gota/dataframe"
	"github.com/go-vgo/robotgo"
	//series "github.com/go-gota/gota/series"
)

func main() {

	// Perhaps the most basic file reading task is slurping a file’s entire contents into memory.
	//
	csvStr, err := ioutil.ReadFile("data.csv")
	check(err)

	// load string data 'dat' to dataframe
	df := dataframe.ReadCSV(strings.NewReader(string(csvStr)))

	// print dataframe 'df2' content
	fmt.Println(df)
	/*
		for i := 0; i < 3; i++ {
			ma := df.Elem(i, 2)
			massnahme := ma.String()
			fmt.Println(massnahme)
			fmt.Println(reflect.TypeOf(massnahme))
			ro := df.Elem(i, 1).String()
			rows_schicht, err := strconv.Atoi(ro)
			if err == nil {
				fmt.Println(rows_schicht)
			}
			fmt.Println(rows_schicht)
			fmt.Println(reflect.TypeOf(rows_schicht))
		}
	*/
	// loop rows
	for i := 0; i < 3; i++ {

		massnahme := df.Elem(i, 2).String()

		if massnahme != "0" {
			// get number of row in Schichtmerkmale
			ro := df.Elem(i, 1).String()
			rows_schicht, err := strconv.Atoi(ro)
			if err == nil {
				fmt.Println(rows_schicht)
			}

			new_nutz(rows_schicht)
			// loop elements
			for j := 2; j < 11; j++ {
				// get elemet
				elem := df.Elem(i, j).String()
				robotgo.KeyTap("tab")
				robotgo.TypeStr(elem)
			}
		}
		next_wo()
	}

	/*
		// insert nutz field
		new_bz()
		fill_bz("AH", "6")
		new_bz()
		fill_bz("EI", "4")
	*/
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
	robotgo.MoveMouse(295, 100)
	//one left click
	robotgo.Click()
}

// make new Nutzung
func new_nutz(rows int) {
	coord := 575 + 20*(rows-1)
	// move mouse to
	robotgo.MoveMouse(222, 506)
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

package main

/*
program for automating TaxMaske
takes as input a csv file named data.csv
and fills data into Taxation Offline Client
version 0.2
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	//"context"
	"io/ioutil"
	// external libreries
	dataframe "github.com/go-gota/gota/dataframe"
	"github.com/go-vgo/robotgo"
	//series "github.com/go-gota/gota/series"
)

func main() {

	// file structue:
	// 0:Abteilung , 1:WE-Typ, 2:number of rows , 3:Umtriebszeit , 4:Waldtyp ,
	// 5: Mart , 6:Afl , 7:LH , 8:NH , 9:D , 10:B , 11:Z , 12:S , 13:R ,
	// 14:FI , 15:TA , 16:LA , 17:KI , 18:DG , 19:TH , 20:BU , 21:EI , 22:HB , 23:AH , 24:ES , 25:ER , 26:RE , 27:SP , 28:EL , 29:SN , 30:SL ,
	// 31:text_1 , 32:text_2 , 33:text_3 , 34:text_4

	// define map[type]type - define dictionary column : baumart
	//FI,TA,LA,KI,BU,EI,AH,ES,ER

	/*var map_ba = map[int]string{
		14: "FI",
		15: "TA",
		16: "LA",
		17: "KI",
		18: "DG",
		19: "TH",
		20: "BU",
		21: "EI",
		22: "HB",
		23: "AH",
		24: "ES",
		25: "ER",
		26: "RE",
		27: "SP",
		28: "EL",
		29: "SN",
		30: "SL",
	}*/

	// load data from csv file
	csvStr, err := ioutil.ReadFile("data/autax_nutz.csv")
	check(err)

	// load string data 'dat' to dataframe
	df := dataframe.ReadCSV(strings.NewReader(string(csvStr)))
	//df_row, _ := df.Dims()
	df_start, err1 := strconv.Atoi(os.Args[1])
	df_end, err2 := strconv.Atoi(os.Args[2])

	if (err1 != nil) && (err2 != nil) {
		fmt.Print("not")
	}
	// print dataframe 'df' content
	fmt.Println(df)
	fmt.Println("working...")

	// loop rows
	for row := df_start; row < df_end+1; row++ {

		we := df.Elem(row, 1).String()

		// if WO than fill in data
		if we == "WO" {

			// *** Maßnahme ***
			massnahme := df.Elem(row, 4).String()
			//fmt.Println("Maßnahme" + massnahme)
			if massnahme != "0" {
				// get number of rows in Schichtmerkmale
				ro := df.Elem(row, 3).String()
				rows_schicht, err := strconv.Atoi(ro)
				if err != nil {
					fmt.Println(rows_schicht)
				}
				// get number of rows in Bestockungsziel
				ro = df.Elem(row, 2).String()
				rows_bz, err := strconv.Atoi(ro)
				if err != nil {
					fmt.Println(rows_bz)
				}

				fmt.Println(row, massnahme)

				new_nutz(rows_schicht, rows_bz)
				// loop columns 2 to 10 (Maßnahmenzeile)
				for col := 4; col < 15; col++ {

					// get elemet
					elem := df.Elem(row, col).String()
					// if column is maßnahmenflaeche
					if col == 6 {
						// replace . by ,
						elem = strings.Replace(elem, ".", ",", 1)
					} else if (col == 14) && (elem == "0") {
						elem = ""
					}
					//fmt.Println("Element: ", elem)

					// fill in data
					//robotgo.KeyTap("tab")
					robotgo.TypeStr(elem)
					robotgo.KeyTap("tab")
					time.Sleep(1 / 2 * time.Second)
				}
			}
		}

		// next Waldort
		next_wo()
		//time.Sleep(250 * time.Millisecond)
	}
	fmt.Println("end")
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

// make new Umtriebszeit
func new_uz() {
	robotgo.MoveMouse(280, 180) // 280, 185
	robotgo.Click()
}

// make new  Waldtyp
func new_wtyp() {
	robotgo.MoveMouse(250, 230) // 250, 240
	robotgo.Click()
}

// make new Nutzung
func new_nutz(rows int, rows_bz int) {

	fmt.Println("Nutzung: ", rows, rows_bz)
	// call offset function
	add := calc_offset(rows, rows_bz)
	// y-coord of add button
	coord := 558 + add + 24*(rows-1) // 575
	// move mouse to
	robotgo.MoveMouse(221, 500+add) // 222, 506
	time.Sleep(1 / 2 * time.Second)
	//one left click
	robotgo.Click()

	time.Sleep(1 / 2 * time.Second)
	fmt.Println("Nutzung ende: ", coord)
	// move mouse to
	robotgo.MoveMouse(224, coord)
	// waiting for "enter" input
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//one left click
	robotgo.Click()
	time.Sleep(1 / 2 * time.Second)
}

// delete old Nutzung
func delete_nutz(rows int, rows_bz int) {
	// call offset function
	add := calc_offset(rows, rows_bz)
	// move mouse to
	robotgo.MoveMouse(221, 608+add+24*(rows-1))
	time.Sleep(1 / 2 * time.Second)
	//one left click
	robotgo.Click()
	// move mouse to
	robotgo.MoveMouse(251, 558+add+24*(rows-1))
	// waiting for "enter" input
	//bufio.NewReader(os.Stdin).ReadBytes('\n')
	// 5 left clicks
	for i := 0; i < 5; i++ {
		robotgo.Click()
	}
	time.Sleep(1 / 2 * time.Second)
}

// calc offset becaues of dynamic GUI
func calc_offset(rows int, rows_bz int) int {
	add := 0
	// if BZ rows more then 2 add offset becaues of dynamic GUI
	if rows_bz == 3 {
		add = 5
	} else if rows_bz > 3 {
		add = 5 + (24 * (rows_bz - 3))
	}
	return add
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

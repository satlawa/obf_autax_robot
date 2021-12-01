package main

/*
program for automating TaxMaske
takes as input a csv file named data.csv
and fills data into Taxation Offline Client
version 0.2
*/

import (
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

	// load data from csv file
	// old "data_BZ_text.csv"
	csvStr, err := ioutil.ReadFile("data/autax_wo.csv")
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

		//if row%0 == 0.0 {
		//}

		we := df.Elem(row, 1).String()

		// if WO than fill in data
		if we == "WO" {

			// *** Text ***
			new_uz()
			time.Sleep(50 * time.Millisecond)
			// loop columns 2 to 11
			for col := 2; col < 11; col++ {

        // get elemet
        elem := df.Elem(row, col).String()

        if (elem != "0") {
          // fill in data
          robotgo.TypeStr(elem)
        }
        robotgo.KeyTap("tab")

        if (col == 5) {
          robotgo.KeyTap("tab")
          robotgo.KeyTap("tab")
          robotgo.KeyTap("tab")
          robotgo.KeyTap("tab")
        } else if (col == 7) {
          new_uelh()
        } else if (col == 9) {
          robotgo.KeyTap("tab")
        }
				time.Sleep(10 * time.Millisecond)

			}
		}
		time.Sleep(30 * time.Millisecond)
		// next Waldort
		next_wo()
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

// make new Umtriebszeit
func new_uelh() {
	robotgo.MoveMouse(355, 230)
	robotgo.Click()
}

// make new  Waldtyp
func new_wtyp() {
	robotgo.MoveMouse(250, 230) // 250, 240
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
	//fmt.Println(coord)
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

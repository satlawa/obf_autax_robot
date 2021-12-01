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

	// define map[type]type - define dictionary column : baumart
	//FI,TA,LA,KI,BU,EI,AH,ES,ER

	var map_ba = map[int]string{
		2:  "FI",
		3:  "TA",
		4:  "LA",
		5:  "KI",
		6:  "SK",
		7:  "ZI",
		8:  "PM",
		9:  "DG",
		10: "KW",
		11: "SF",
		12: "FO",
		13: "TH",
		14: "FB",
		15: "AC",
		16: "AG",
		17: "AZ",
		18: "EB",
		19: "FZ",
		20: "GK",
		21: "HT",
		22: "JL",
		23: "CJ",
		24: "KK",
		25: "KO",
		26: "AN",
		27: "AB",
		28: "CH",
		29: "PU",
		30: "SN",
		31: "BU",
		32: "EI",
		33: "HB",
		34: "AH",
		35: "SA",
		36: "FA",
		37: "EA",
		38: "ES",
		39: "UL",
		40: "QP",
		41: "QR",
		42: "EZ",
		43: "RE",
		44: "FE",
		45: "ER",
		46: "GE",
		47: "AV",
		48: "KB",
		49: "TK",
		50: "WO",
		51: "SG",
		52: "NU",
		53: "JN",
		54: "LI",
		55: "LS",
		56: "LW",
		57: "BI",
		58: "PO",
		59: "AS",
		60: "WP",
		61: "SP",
		62: "HP",
		63: "WD",
		64: "SW",
		65: "EK",
		66: "RK",
		67: "EE",
		68: "EL",
		69: "ME",
		70: "RO",
		71: "TB",
		72: "GB",
		73: "ST",
		74: "SL",
		75: "BL",
	}

	// load data from csv file
	csvStr, err := ioutil.ReadFile("data/autax_bz.csv")
	check(err)

	// load string data 'dat' to dataframe
	df := dataframe.ReadCSV(strings.NewReader(string(csvStr)))
	//df_row, _ := df.Dims()
	df_start, err1 := strconv.Atoi(os.Args[1])
	df_end, err2 := strconv.Atoi(os.Args[2])
	awk := os.Args[3]

	if (err1 != nil) && (err2 != nil) {
		fmt.Print("not")
	}
	// print dataframe 'df' content
	fmt.Println(df)
	fmt.Println("working...")

	// create Auswertekategorie max width to avoid trubles
	if awk == "y" {
		create_awk()
	}

	// loop rows
	for row := df_start; row < df_end+1; row++ {

		//if row%10 == 0.0 {
		//}

		we := df.Elem(row, 1).String()

		// if WO than fill in data
		if we == "WO" {
			fmt.Println("------------")
			fmt.Println(we)
			fmt.Println(row)
			// *** Bestockungsziel ***
			//fmt.Println("Bestockungsziel")
			// delete old
			delete_bz()
			time.Sleep(25 * time.Millisecond)
			// loop though columns 11 to 19 (BA-Bestockungsziele)
			for col := 2; col < 76; col++ {
				// get elemet
				elem := df.Elem(row, col).String()
				if elem != "0" {
					// get baumart from dictionary
					elem_ba := map_ba[col]
					// insert nutz field and fill in data
					new_bz()
					time.Sleep(50 * time.Millisecond)
					fill_bz(elem_ba, elem)
				}
			}
			robotgo.KeyTap("enter")
			time.Sleep(50 * time.Millisecond)
		}

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

// create max width Auswertekategorie
func create_awk() {
	// move mouse to
	robotgo.MoveMouse(1712, 156)
	//one left click
	robotgo.Click()
	robotgo.TypeStr("SO")
	robotgo.KeyTap("tab")
	robotgo.TypeStr("51197")
	robotgo.KeyTap("tab")
	robotgo.TypeStr("0")
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
	robotgo.MoveMouse(1340, 340) // 1470, 340  // 1447, 340  // 822, 300
	//one left click
	robotgo.Click()
	// move mouse to
	robotgo.MoveMouse(1370, 290) // 1500, 290  // 1475, 290  // 822, 300
	// 5 left clicks
	for i := 0; i < 6; i++ {
		robotgo.Click()
	}
}

// make new Bestockungsziel
func new_bz() {
	// move mouse to
	robotgo.MoveMouse(1340, 290) // 1470, 290  // 1450, 290  // 822, 300
	//one left click
	robotgo.Click()
}

// fill Bestockungsziel with data
func fill_bz(ba_art string, ba_anteil string) {

	robotgo.TypeStr(ba_art)
	robotgo.KeyTap("tab")
	robotgo.TypeStr(ba_anteil)
	time.Sleep(1 / 10 * time.Second)
}

func new_text() {
	// move mouse to
	robotgo.MoveMouse(350, 350)
	//one left click
	robotgo.Click()
	robotgo.KeyTap("a", "ctrl")
	robotgo.KeyTap("delete")
}

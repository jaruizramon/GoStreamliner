package actions

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// GLOBALS START
var session []string
var is_quit = false
var user32_dll = windows.NewLazyDLL("user32.dll")
var GetKeyState = user32_dll.NewProc("GetKeyState")
var coordX, coordY int64
var coordsX, coordsY []int64
var dts []float64

// GLOBALS END

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// var readKeyPressBuff = bufio.NewReader(os.Stdin)
//
//	func CheckQuit(input chan rune) {
//		char, _, err := readKeyPressBuff.ReadRune()
//		if err != nil {
//			log.Fatal(err)
//		}
//		input <- char
//	}
func wasESCPressed() bool {
	r1, _, _ := GetKeyState.Call(27) // Call API to get ESC key state.
	return r1 == 65409               // Code for KEY_UP event of ESC key.
}
func GetAdbCoords() {

	s := "cd adb"
	args := strings.Split(s, " ")

	changeDir := exec.Command(args[0], args[1:]...)
	time.Sleep(1 * time.Second)
	fmt.Println(changeDir.Stdout)

	s2 := "adb shell getevent"
	args2 := strings.Split(s2, " ")
	var isQuit = false

	cmd := exec.Command(args2[0], args2[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("log ->", err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)
	var reinitTimer = true
	start := time.Now() // declaration

	for isQuit == false {
		if wasESCPressed() {
			isQuit = true
			break
		}
		time.Sleep(time.Millisecond * 100)
		if reinitTimer == true {
			start = time.Now()
			reinitTimer = false
		}
		line, _, _ := buf.ReadLine()
		if line == nil {
			break
		}
		var appendableString, secondFourDigits, lastEightDigits string

		s := string(line)
		if len(s) >= 37 {
			appendableString = s[len(s)-18:]
			// #### #### ########
			secondFourDigits, lastEightDigits = appendableString[5:9], appendableString[10:18]
			if secondFourDigits == "0035" { // X coord
				coordX, err = strconv.ParseInt(lastEightDigits, 16, 64)
				if err != nil {
					panic(err)
				}
			} else if secondFourDigits == "0036" { // Y coord
				coordY, err = strconv.ParseInt(lastEightDigits, 16, 64)
				if err != nil {
					panic(err)
				}
			}
		} else {
			appendableString = "loading..."
		}
		// PROBLEM START
		var duration float64
		if coordX != 0 {
			coordsX = append(coordsX, coordX)
			fmt.Printf("\n(%d ,", coordX)
		}
		if coordY != 0 {
			coordsY = append(coordsY, coordY)
			duration = float64(time.Since(start).Seconds())
			dts = append(dts, duration)
			reinitTimer = true
			fmt.Printf(" %d) ---> %.2fs", coordY, duration)
		}
		// PROBLEM END
		time.Sleep(1 * time.Millisecond)
	} // for loop
	coordsX = append([]int64{0}, coordsX...) // ADD ELEMENT FROM idx 0
	coordsX = coordsX[:len(coordsX)-1]       // REMOVE LAST ELEMENT
	coordsX = coordsX[:len(coordsX)-1]       // REMOVE LAST ELEMENT

	dts = append([]float64{0}, dts...) // ADD ELEMENT FROM idx 0
	dts = dts[:len(dts)-1]             // REMOVE LAST ELEMENT

	fmt.Println("\n\nConsole end")

	file, err := os.OpenFile("testcase.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	filePopulator := bufio.NewWriter(file)

	fmt.Printf("%d %d %d", len(coordsX), len(coordsY), len(dts))

	var session []string

	if len(coordsX) == len(coordsY) && len(coordsX) == len(dts) {
		for i := range coordsX {
			if dts[i] > 0.1 {
				s = fmt.Sprintf("%d,%d,%.2f\n", coordsX[i], coordsY[i], dts[i])
				session = append(session, s)
			}
		}
	}
	session = removeDuplicateStr(session)
	for i := range session {
		filePopulator.WriteString(session[i])
	}
	filePopulator.Flush()
	file.Close()
} // func

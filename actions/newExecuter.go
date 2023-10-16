package actions

// read this shit to improve this shit
// https://android.stackexchange.com/questions/236037/faster-alternative-to-input

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)
func ExecuteShell() {

	var (
		x string
		y string
		dt    string
		shellString strings.Builder
		swipe string
	)
	// open file
	f, err := os.Open("testy.csv")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()
	commands := [3]string{"adb","shell \"","input tap 150 40 && "}

	os.Chdir("adb")
	cwd, _ := os.Getwd()

	var executableString string

	csvReader := csv.NewReader(f)
	for {
		action, err := csvReader.Read()
		if err == io.EOF || action == nil {
			closeQuote := shellString.String()
			closeQuote = closeQuote[:len(shellString.String())-3] + "\""
			executableString = closeQuote
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		x = action[0]
		y = action[1]
		dt = action[2]
		swipe = action[3]

		fmt.Printf("%s %s %s %s \n", x,y,dt,swipe)

		var swiper string

		if len(swipe) <= 52{
			swiper = " "
		} else {
			swiper = string([]rune(swipe)[52:])
		}
// PROBLEM START
		var commandy string
		if x != "-1" && y != "-1" { // if tap!
			fmt.Printf("%s %s %s %s \n", x,y,dt,swipe)
			shellString.WriteString(fmt.Sprintf("echo '%s %s %s' && ", x, y, dt))
			commandy = fmt.Sprintf("input tap %s %s && sleep %s && \n", x, y,dt)
			shellString.WriteString(commandy)
		} else { // it's a swipe!
			fmt.Printf("%s %s %s %s \n", x,y,dt,swipe)
			shellString.WriteString(fmt.Sprintf("echo 'SWIPE %s' && ", swiper))
			commandy = fmt.Sprintf("input swipe %s && sleep %s \n" , swiper, "1.5")
			shellString.WriteString(commandy)
			
		}
	}
	cmd := exec.Command(cwd, commands[1], executableString)
	fmt.Printf("adb %s%s", commands[1], executableString)
// PROBLEM
	cmd.Run()
}
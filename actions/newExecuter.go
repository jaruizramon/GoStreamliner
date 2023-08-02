package actions

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
		// dt    string
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
	commands := [2]string{"adb","shell \""}

	os.Chdir("adb")
	cwd, _ := os.Getwd()

	var executableString string

	csvReader := csv.NewReader(f)
	for {
		action, err := csvReader.Read()
		if err == io.EOF || action == nil {
			closeQuote := shellString.String()
			closeQuote = closeQuote[:len(shellString.String())-3]
			closeQuote = closeQuote + "\""
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



		if x != "-1" && y != "-1" { // if swipe
			shellString.WriteString(fmt.Sprintf("input tap %s %s && sleep %s &&\n", x, y,dt))
		} else { // it's a tap!
			shellString.WriteString(fmt.Sprintf("input swipe %s & sleep %s &&\n", swipe, dt))
		}

	}
	fmt.Println(executableString)
	cmd := exec.Command(cwd, commands[0], commands[1], executableString)
	fmt.Printf("%s %s %s %s", cwd, commands[0], commands[1], executableString)
	cmd.Run()


	fmt.Println("END.")

	
}

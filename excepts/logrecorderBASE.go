package excepts

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetAdbCoordss() {
	s := "cd adb"
	args := strings.Split(s, " ")

	changeDir := exec.Command(args[0], args[1:]...)
	time.Sleep(1 * time.Second)
	fmt.Println(changeDir.Stdout)

	s2 := "adb shell getevent"
	args2 := strings.Split(s2, " ")
	var is_quit bool = false

	cmd := exec.Command(args2[0], args2[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("log ->", err)
	}
	cmd.Start()
	buf := bufio.NewReader(stdout)

	for is_quit == false {
		start := time.Now()
		line, _, _ := buf.ReadLine()
		if line == nil {
			break
		}
		var appendableString, secondFourDigits, lastEightDigits string
		var coordX, coordY int64
		var coordsX, coordsY []int64

		s := string(line)
		if len(s) >= 37 {
			appendableString = s[len(s)-18:]
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
		var duration float64
		if coordX != 0 {

			duration = float64(time.Since(start))
			coordsX = append(coordsX, coordX)
			fmt.Printf("\n(%d ,", coordX)
		}
		if coordY != 0 {
			coordsY = append(coordsY, coordY)
			fmt.Printf(" %d) -> %.2f s", coordY, duration)
		}
		time.Sleep(16 * time.Millisecond)
	}
	fmt.Println("Console end")

}

package excepts

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var session []string
var is_quit = false

//var readKeyPressBuff = bufio.NewReader(os.Stdin)
//
//func CheckQuit(input chan rune) {
//	char, _, err := readKeyPressBuff.ReadRune()
//	if err != nil {
//		log.Fatal(err)
//	}
//	input <- char
//}

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

	ch := make(chan string)
	go func(ch chan string) {
		// disable input buffering
		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
		// do not display entered characters on the screen
		exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
		var b []byte = make([]byte, 1)

		ch <- string(b)

	}(ch) // pasted this goroutine, gotta read about them channel directions

	for isQuit == false {

		select {
		case stdin, _ := <-ch:
			if stdin == "q" {
				isQuit = true
				fmt.Printf(stdin)
			}

		default:
			isQuit = false // assert to stay in loop
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
		var coordX, coordY int64
		var coordsX, coordsY []int64

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
		var duration float64
		if coordX != 0 {
			coordsX = append(coordsX, coordX)
			fmt.Printf("\n(%d ,", coordX)
		}
		if coordY != 0 {
			coordsY = append(coordsY, coordY)
			duration = float64(time.Since(start).Seconds())
			reinitTimer = true
			fmt.Printf(" %d) -> %.2fs", coordY, duration)

			if len(coordsX) == 0 { //
			} else {
				s = fmt.Sprintf("%d,%d,%.2f", coordsX[len(coordsX)-1], coordY, duration)
			}
		}
		time.Sleep(10 * time.Millisecond)

		session = append(session, s)

	} // for loop

	fmt.Println("Console end")

	file, err := os.OpenFile("testcase.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	filePopulator := bufio.NewWriter(file)

	for _, data := range session {
		_, _ = filePopulator.WriteString(data + "\n")
	}

	filePopulator.Flush()
	file.Close()
} // func

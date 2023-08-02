package actions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/windows"
)

// vars START

var (
	session          []string
	is_quit          = false
	user32_dll       = windows.NewLazyDLL("user32.dll")
	GetKeyState      = user32_dll.NewProc("GetKeyState")
	coordX, coordY   int64
	coordsX, coordsY []int64
	dts              []float64
	swipes           []string
	start            = time.Now() // declaration
	reinitTimer      = true
)

// vars END

func GetAdbCoords() byte {

	s := "cd adb"
	args := strings.Split(s, " ")

	cmd := exec.Command(args[0], args[1:]...)
	time.Sleep(3 * time.Second)
	fmt.Println(cmd.Stdout)

	s = "adb shell getevent"
	args2 := strings.Split(s, " ")
	var isQuit = false

	cmd = exec.Command(args2[0], args2[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal("log ->", err)
	}
	cmd.Start()


	
	buf := bufio.NewReader(stdout)
	reinitTimer = true

	for isQuit == false {
		if wasESCPressed() {
			isQuit = true
			break
		}

		time.Sleep(time.Millisecond * 1)
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
		var estrin strings.Builder
		var dt float64
		if coordX != 0 {
			//coordsX = append(coordsX, coordX)
			estrin.WriteString(fmt.Sprintf("%v,", coordX))
			//fmt.Printf("\n(%d ,", coordX)
			if coordY != 0 {
				//coordsY = append(coordsY, coordY)
				estrin.WriteString(fmt.Sprintf("%v,", coordY))
				dt = time.Since(start).Seconds()
				//dts = append(dts, duration)
				estrin.WriteString(fmt.Sprintf("%.2f,\n", dt)) // %.2f
				if dt < 0.011 == false {
					fmt.Print(estrin.String())
					session = append(session, estrin.String())
				}
				reinitTimer = true
			}
		}

		estrin.Reset()

		// PROBLEM END
		time.Sleep(10 * time.Millisecond)
	} // for loop
	// coordsX = coordsX[:len(coordsX)-1] // REMOVE LAST ELEMENT
	// dts = append([]float64{1}, dts...) // ADD ELEMENT FROM idx 0
	// dts = dts[:len(dts)-1]             // REMOVE LAST ELEMENT

	fmt.Println("\n\nConsole end")

	file, err := os.OpenFile("testy.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	filePopulator := bufio.NewWriter(file)

	fmt.Printf("%d %d %d", len(coordsX), len(coordsY), len(dts))

	session = removeDuplicateStr(session)
	for i := range session {
		filePopulator.WriteString(session[i])
	}
	filePopulator.Flush()
	file.Close()

	return 0
} // func

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

func wasESCPressed() bool {
	r1, _, _ := GetKeyState.Call(27) // Call API to get ESC key state.
	return r1 == 65409               // Code for KEY_UP event of ESC key.
}

func Swipe(direction string) {

	var param string
	var estrin strings.Builder
	dt := time.Since(start).Seconds()

	switch direction {
	case "up":
		param = "690 400 690 450 100"
	case "down":
		param = "690 450 690 400 100"
	case "left":
		param = "690 450 690 400 100" // tbc
	case "right":
		param = "690 450 690 400 100" // tbc
	}

	uniqueTime := time.Now()
	fmtTime := uniqueTime.Format(time.UnixDate)

	appendableSwipeString := fmt.Sprintf("%s->adb shell input swipe %s", fmtTime, param)
	fmt.Println(appendableSwipeString)

	swipes = append(swipes, appendableSwipeString)
	swiper := string([]rune(appendableSwipeString)[30:])
	// coordsX = append(coordsX, -1)
	// coordsY = append(coordsY, -1

	estrin.WriteString(fmt.Sprintf("%v,%v,%v,%v", -1, -1, dt, appendableSwipeString))
	appndS := strings.Split(swiper, " ")
	session = append(session, estrin.String()+"\n")
	sp := exec.Command(appndS[0], appndS[1:]...)

	sp.Run()

}

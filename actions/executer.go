package actions

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	x     string
	y     string
	dt    string
	swipe string
	comm  = exec.Command(" ")
)

func Execute() {
	comm.Start()

	// open file
	f, err := os.Open("testy.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	os.Chdir("adb")
	cwd, _ := os.Getwd()
	csvReader := csv.NewReader(f)
	for {
		action, err := csvReader.Read()
		if err == io.EOF || action == nil {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		x = action[0]
		y = action[1]
		dt = action[2]
		swipe = action[3]
		// do something with read line
		fmt.Printf("%+v\n", action)

		if x == "-1" && y == "-1" {

			fmt.Println(len(swipe))
			swiper := string([]rune(swipe)[30:])
			fmt.Println(swipe)
			swipeArg := strings.Split(swiper, " ")
			comm = exec.Command(swipeArg[0], swipeArg[1:]...)
			comm.Run()
			fmt.Println(cwd, swipeArg[1], swipeArg[2], swipeArg[3], swipeArg[4], swipeArg[5], swipeArg[6], swipeArg[7], swipeArg[8])

		} else if swipe == " " || swipe == "" {
			comm = exec.Command(cwd+"\\adb.exe", "shell", "input", "tap", x, y)
			comm.Run()
			fmt.Println(fmt.Sprintf("{ x -> %s | y -> %s | Zzz -> %s}", x, y, dt))

			//fmt.Println(cwd+"\\adb.exe", "shell", "input", "tap", x, y)

		}
		sleep, _ := strconv.ParseFloat(dt, 64)
		time.Sleep(time.Duration(sleep) * time.Second)
	}

	fmt.Println("END.")

}

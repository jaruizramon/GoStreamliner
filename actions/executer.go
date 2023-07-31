package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
)

func Execute() {

	var (
		x  string
		y  string
		dt string
	)

	rawContent, _ := ioutil.ReadFile("testcase.csv")
	ioContent := strings.NewReader(string(rawContent))
	df := dataframe.ReadCSV(ioContent)
	comm := exec.Command(" ")

	os.Chdir("adb")

	var idx = 0
	for idx <= df.Nrow()-1 {

		x = df.Elem(idx, 0).String()
		y = df.Elem(idx, 1).String()
		dt = df.Elem(idx, 2).String()

		dtFloat, _ := strconv.ParseFloat(dt, 32)
		cwd, _ := os.Getwd()

		if x == "-1" {
			swipe := df.Elem(idx, 3).String()

			swiper := strings.Split(swipe, " ")
			swiper = append([]string{cwd + "\\adb.exe"}, swiper...)
			comm = exec.Command(swiper[0], swiper[1:]...)
			fmt.Println(swiper)

		} else {

			comm = exec.Command(cwd+"\\adb.exe", "shell", "input", "tap", x, y)
			fmt.Println(fmt.Sprintf("{ x -> %s | y -> %s | Zzz -> %s}", x, y, dt))
			//fmt.Println(cwd+"\\adb.exe", "shell", "input", "tap", x, y)
			comm.Run()
		}

		//fmt.Println(cwd)

		time.Sleep(time.Duration(dtFloat) * time.Second)
		idx += 1

	}

	idx = 0
	fmt.Println("EXECUTE END.")
}

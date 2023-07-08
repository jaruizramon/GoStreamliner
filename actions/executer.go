package actions

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
	for idx != df.Nrow() {

		x = df.Elem(idx, 0).String()
		y = df.Elem(idx, 1).String()
		dt = df.Elem(idx, 2).String()

		//xInt, _ := strconv.ParseInt(x, 10, 32)
		//yInt, _ := strconv.ParseInt(y, 10, 32)
		dtFloat, _ := strconv.ParseFloat(dt, 32)

		//xs = append(xs,xInt)
		//ys = append(ys,yInt)
		//dts = append(dts,dtFloat)

		fmt.Println(fmt.Sprintf("{ x -> %s | y -> %s | Zzz -> %s}", x, y, dt))

		//subCommand := fmt.Sprintf("adb.exe shell input tap %s %s\n", x , y)

		//subCommandArgs := strings.Split(subCommand, " ")
		//exec.Command(subCommandArgs[0], subCommandArgs[1:]...)
		cwd, _ := os.Getwd()
		comm = exec.Command(cwd+"\\adb.exe", "shell", "input", "tap", x, y)
		fmt.Println(cwd+"\\adb.exe", "shell", "input", "tap", x, y)
		comm.Run()

		fmt.Println(cwd)

		time.Sleep(time.Duration(dtFloat) * time.Second)
		idx += 1

	}

}

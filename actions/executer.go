package actions

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"os/exec"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func Execute(){
	var (
		x string
		y string
		dt string
	)

	ou := exec.Command("cd","..")
	ou1 := exec.Command("cd", "adb")

	fmt.Println(ou , ou1)

	time.Sleep(1 * time.Second)

	rawContent, _ := ioutil.ReadFile("testcase.csv")
	ioContent := strings.NewReader(string(rawContent))
	df := dataframe.ReadCSV(ioContent)

	var idx = 0
	for idx != df.Nrow(){

		x = df.Elem(idx, 0).String()
		y = df.Elem(idx, 1).String()
		dt = df.Elem(idx, 2).String()

		//xInt, _ := strconv.ParseInt(x, 10, 32)
		//yInt, _ := strconv.ParseInt(y, 10, 32)
		dtFloat, _ := strconv.ParseFloat(dt, 32)
		//
		//xs = append(xs,xInt)
		//ys = append(ys,yInt)
		//dts = append(dts,dtFloat)

		fmt.Println(fmt.Sprintf("{ x -> %s | y -> %s | Zzz -> %s}", x, y, dt))

		subCommand := fmt.Sprintf("adb shell input tap %s %s", x , y)
		subCommandArgs := strings.Split(subCommand, " ")
		ou3 := exec.Command(subCommandArgs[0], subCommandArgs[1:]...)
		fmt.Println(ou3)
		time.Sleep(time.Duration(dtFloat) * time.Second)
		idx += 1

	}
	// Start executing commands here

}

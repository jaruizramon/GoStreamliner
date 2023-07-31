package main

import (
	"fmt"

	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"

	"Gostreamliner/actions"
)

func main() {

	wnd := nucular.NewMasterWindow(0, "ADB RECORDER LITE!", updatefn)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 2.0))
	wnd.Main()

}

func updatefn(w *nucular.Window) {
	w.RowScaled(50).Dynamic(1)

	if w.ButtonText("RECORD ADB INPUT") {
		go actions.GetAdbCoords()

	} else if w.ButtonText("EXECUTE RECORDED ADB INPUT") {
		go actions.Execute()

	} else if w.ButtonText("EXECUTE ALL TEST CASES") {
		fmt.Println("NOT AVAILABLE NOW!")
	} else if w.ButtonText("SWIPE DOWN") {
		go actions.SwipeDown()
	}

}

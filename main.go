package main

import (
	"Gostreamliner/actions"
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"
)

func main() {

	wnd := nucular.NewMasterWindow(0, "ADB RECORDER LITE!", updatefn)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 2.0))
	wnd.Main()

}

func updatefn(w *nucular.Window) {
	w.Row(50).Dynamic(1)
	if w.ButtonText("RECORD ADB INPUT") {
		actions.GetAdbCoords()
	} else if w.ButtonText("EXECUTE RECORDED ADB INPUT") {
		actions.Execute()
	} else if w.ButtonText("EXECUTE ALL TEST CASES") {
		actions.Execute()
	}

}

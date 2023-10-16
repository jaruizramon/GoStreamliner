package main

import (
	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/style"

	"Gostreamliner/actions"
)

// global state var
var (
	globalWaiter = [7] bool{
		true, // Recording Process ON Done
		true, // Recording Process OFF Done
		true, // Printing Recorded Session Done
		true, // Swiping up Done
		true, // Swiping down Done
		true, // Swiping left Done
		true, // Swiping right Done
	}
) 

func main() {

	wnd := nucular.NewMasterWindow(0, "ADB RECORDER LITE!", updatefn)
	wnd.SetStyle(style.FromTheme(style.DarkTheme, 2.0))
	wnd.Main()
}

func updatefn(w *nucular.Window) {
	w.RowScaled(50).Dynamic(1)

	if w.ButtonText("RECORD ADB INPUT") { 
		go actions.GetAdbCoords(&globalWaiter[0])
		// 0
		for !globalWaiter[0]{}
		
	} else if w.ButtonText("Stop Recording Input"){
		go actions.QuitRecording(&globalWaiter[0], &globalWaiter[1])
		//  
		for !globalWaiter[1]{}
	} else if w.ButtonText("Print out recorded ADB session.") { 
		go actions.ExecuteShell()
		// 2
		for !globalWaiter[2]{}
	} else if w.ButtonText("UP") {
		go actions.Swipe("down")
		// 3
		for !globalWaiter[3]{}
	} else if w.ButtonText("UP") {
		go actions.Swipe("up")
		// 4
		for !globalWaiter[4]{}
	} else if w.ButtonText("LEFT") {
		go actions.Swipe("left")
		// 5
		for !globalWaiter[5]{}
	} else if w.ButtonText("RIGHT") {
		go actions.Swipe("right")
		// 6 
		for !globalWaiter[6]{}
	}

}

package main

import (
	"fyne.io/fyne/v2/app"
	"game/screens"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Tom and Jerry game")
	w.CenterOnScreen()
	screens.NewMainScene(w)
	w.CenterOnScreen()
	w.ShowAndRun()
}

package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GoChat")

	w.SetContent(widget.NewLabel("This better be a good platform..."))
	w.Resize(fyne.NewSize(750, 200))

	w.ShowAndRun()
}

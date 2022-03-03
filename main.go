// All shit has this. Guessing it's like a namespace or something idk.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Declares a new process and a window for it
	a := app.New()
	w := a.NewWindow("GoChat")

	// Puts some words in the window, so you can know that it is doing anything and resizes it.
	w.SetContent(widget.NewLabel("This better be a good platform..."))
	w.Resize(fyne.NewSize(750, 200))

	// This could be separately w.Show() and a.Run(), but you can shorthand it if you don't
	// need to do anything else after showing whichever window.
	w.ShowAndRun()
}

/*
TODO:
1) Swap functionality of Enter and Shift + Enter for widget.NewMultiLineEntry(). Currently 'Enter'
   starts a new line and 'Shift + Enter' submits.
2) Move shit to the bottom of the window.
3) Figure out how to move shit just in general
*/

// All shit has this. Guessing it's like a namespace or something idk.
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	out := widget.NewLabel("Enter some text.")
	in := widget.NewEntry()

	message := ""

	// Stores characters as they are tyed in the 'message' variable. Probably mostly uneccesary
	// pretty much just there so the send button works, and that was mostly just as a test.
	in.OnChanged = func(content string) {
		message = content
	}

	// Changes label to be whatever was in the entry field when the 'Enter' key is pressed then clears field.
	in.OnSubmitted = func(content string) {
		if content != "" {
			out.SetText(content)
			in.SetText("")
		}
	}

	send := widget.NewButton("Send", func() {
		if message != "" {
			out.SetText(message)
			in.SetText("")
		}
	})

	return out, in, send
}

func main() {
	// Declares a new process and a window for it
	a := app.New()
	w := a.NewWindow("GoChat")

	// Puts some words in the window, so you can know that it is doing anything and resizes it.
	// w.SetContent(widget.NewLabel("This better be a good platform..."))

	w.SetContent(container.NewVBox(makeUI()))

	w.Resize(fyne.NewSize(900, 500))

	// This could be separately w.Show() and a.Run(), but you can shorthand it if you don't
	// need to do anything else after showing whichever window.
	w.ShowAndRun()
}

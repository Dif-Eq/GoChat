/*
TODO:
1) Swap functionality of Enter and Shift + Enter for widget.NewMultiLineEntry(). Currently 'Enter'
   starts a new line and 'Shift + Enter' submits.
2) Tie window element locations to the size of the window, whatever that is at the time.
3) Should prob get rid of the global vars since that's not a great way to do that.
4) There has GOT to be a way to obtain the dimensions of the window at any time. Figure that out.
5) Instead of rewriting a label, have the sent text saved on screen in some container
6) Add some of those sick games like iMessage has that you can play with people
*/

// All shit has this. Guessing it's like a namespace or something idk.
package main

import (
	"GoChat/Source/api"
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// GLOBAL VARIABLES
var WINDOW_X float32 = 900
var WINDOW_Y float32 = 500

func formatTimeString(timeValue time.Time) string {
	return timeValue.Local().Format(time.RFC822)
}

func renderMessages(messages []api.Message, widget *widget.Label) {
	widget.SetText("")
	for _, message := range messages {
		fmt.Print(message)
		widget.SetText(widget.Text + fmt.Sprintf("\n%v (%v): %v",
			message.Username,
			formatTimeString(message.CreatedAt),
			message.Contents))
	}
}

func makeUI(messages []api.Message) (*widget.Label, *widget.Entry, *widget.Button) {
	out := widget.NewLabel("Enter some text.")
	in := widget.NewEntry()

	message := ""

	renderMessages(messages, out)

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

			message := api.Message{
				Contents: message,
			}

			api.CreateMessage(message)
			renderMessages(api.GetMessages(), out)
			in.SetText("")
		}
	})

	// Finally figured out how to manipulate widgets. Can use .Resize() and .Move() if widgets are in a
	// container without a predetermined layout as in container.NewWithoutLayout().
	// 37 for the Y component of fyne.NewSize() is the smallest value that didn't have a weird scoll bar.
	var entryFieldX float32 = WINDOW_X * 0.9
	var entryFieldY float32 = 37
	var buttonX float32 = 60
	var buttonY float32 = entryFieldY

	in.Resize(fyne.NewSize(entryFieldX, entryFieldY))
	in.Move(fyne.NewPos(9, WINDOW_Y-46))

	send.Resize(fyne.NewSize(buttonX, buttonY))
	send.Move(fyne.NewPos(entryFieldX+15, WINDOW_Y-46))

	return out, in, send
}

func main() {
	// Declares a new process and a window for it
	a := app.New()
	w := a.NewWindow("GoChat")

	messages := api.GetMessages()

	// Creates a container with no pre-made layout for the makeUI() function allowing for control over
	// size and location of widgets created there.
	w.SetContent(container.NewWithoutLayout(makeUI(messages)))

	w.Resize(fyne.NewSize(WINDOW_X, WINDOW_Y))

	// This could be separately w.Show() and a.Run(), but you can shorthand it if you don't
	// need to do anything else after showing whichever window.
	w.ShowAndRun()
}

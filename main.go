package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Create a new application with a unique ID
	myApp := app.NewWithID("whatever")
	myWindow := myApp.NewWindow("PDF Tools")

	//optimize form
	optimizeForm := InitOptimizeForm(myWindow)
	// Create the splitting form
	splitForm := InitSplitForm(myWindow)

	// Combine both forms in a vertical box
	content := container.NewVBox(optimizeForm, widget.NewSeparator(), splitForm)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(1400, 900))
	myWindow.ShowAndRun()
}

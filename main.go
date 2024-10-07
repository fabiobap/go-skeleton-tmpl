package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	form1 := container.NewVBox(optimizeForm)
	form2 := container.NewVBox(layout.NewSpacer(), splitForm)

	tabs := container.NewAppTabs(
		container.NewTabItem("Optimize a PDF file", container.NewPadded(form1)),
		container.NewTabItem("Split a PDF file", form2),
	)

	myWindow.SetContent(tabs)

	myWindow.Resize(fyne.NewSize(1000, 600))
	myWindow.ShowAndRun()
}

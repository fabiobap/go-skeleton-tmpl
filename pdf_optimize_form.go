package main

import (
	"log"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func InitOptimizeForm(myWindow fyne.Window) *widget.Form {
	// Create the optimization form
	optimizeFileEntry := widget.NewEntry()
	optimizeFileEntry.SetPlaceHolder("No file selected")

	optimizeUploadButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					dialog.ShowInformation("Error", "Failed to open file!", myWindow)
					log.Println("Failed to open file:", err)
					return
				}
				if reader == nil {
					return
				}
				optimizeFileEntry.SetText(reader.URI().Path())
			}, myWindow)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".pdf"}))

		location, err := GetExecutableDirectory()
		if err != nil {
			dialog.ShowInformation("Error", "Failed to get executable directory!", myWindow)
			log.Println("Failed to get executable directory:", err)
			return
		}
		fileDialog.SetLocation(location)

		fileDialog.Show()
	})

	optimizeForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "PDF File", Widget: container.NewVBox(layout.NewSpacer(), optimizeFileEntry, optimizeUploadButton)},
		},
		OnSubmit: func() {
			filePath := optimizeFileEntry.Text
			if filePath == "" {
				dialog.ShowInformation("Error", "No file selected!", myWindow)
				log.Println("No file selected")
				return
			}

			dir, file := filepath.Split(filePath)
			compressedFilePath := filepath.Join(dir, "compressed_"+file)

			err := api.OptimizeFile(filePath, compressedFilePath, nil)
			if err != nil {
				dialog.ShowInformation("Error", "Failed to compress file", myWindow)
				log.Println("Failed to compress file:", err)
				return
			}
			dialog.ShowInformation("Success", "File compressed successfully!", myWindow)
			log.Println("File compressed successfully:", compressedFilePath)
		},
	}

	return optimizeForm
}

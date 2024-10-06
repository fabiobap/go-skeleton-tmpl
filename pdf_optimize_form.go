package main

import (
	"log"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func InitOptimizeForm(myWindow fyne.Window) *widget.Form {
	// Create the optimization form
	optimizeLabel := widget.NewLabel("Optimize a PDF file")
	optimizeFileEntry := widget.NewEntry()
	optimizeFileEntry.SetPlaceHolder("No file selected")
	optimizeFileEntry.Resize(fyne.NewSize(300, optimizeFileEntry.MinSize().Height))

	optimizeUploadButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
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
			log.Println("Failed to get executable directory:", err)
			return
		}
		fileDialog.SetLocation(location)

		fileDialog.Show()
	})

	optimizeForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "PDF File", Widget: container.NewVBox(optimizeLabel, container.NewHBox(optimizeFileEntry, optimizeUploadButton))},
		},
		OnSubmit: func() {
			filePath := optimizeFileEntry.Text
			if filePath == "" {
				log.Println("No file selected")
				return
			}

			dir, file := filepath.Split(filePath)
			compressedFilePath := filepath.Join(dir, "compressed_"+file)

			err := api.OptimizeFile(filePath, compressedFilePath, nil)
			if err != nil {
				log.Println("Failed to compress file:", err)
				return
			}

			log.Println("File compressed successfully:", compressedFilePath)
		},
	}

	return optimizeForm
}

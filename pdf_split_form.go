package main

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func InitSplitForm(myWindow fyne.Window) *widget.Form {
	splitFileEntry := widget.NewEntry()
	splitFileEntry.SetPlaceHolder("No file selected")
	splitFileEntry.Resize(fyne.NewSize(300, splitFileEntry.MinSize().Height))

	splitUploadButton := widget.NewButton("Browse", func() {
		fileDialog := dialog.NewFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					log.Println("Failed to open file:", err)
					return
				}
				if reader == nil {
					return
				}
				splitFileEntry.SetText(reader.URI().Path())
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

	pagesEntry := widget.NewEntry()
	pagesEntry.SetPlaceHolder("Number of pages per split")

	splitForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "PDF File", Widget: container.NewVBox(container.NewHBox(splitFileEntry, splitUploadButton))},
			{Text: "Pages", Widget: pagesEntry},
		},
		OnSubmit: func() {
			filePath := splitFileEntry.Text
			if filePath == "" {
				log.Println("No file selected")
				return
			}

			pages, err := strconv.Atoi(pagesEntry.Text)
			if err != nil || pages <= 0 {
				log.Println("Invalid number of pages")
				return
			}

			err = splitPDF(filePath, pages)
			if err != nil {
				log.Println("Failed to split file:", err)
				return
			}
		},
	}

	return splitForm
}

// splitPDF splits the PDF file into smaller files with the specified number of pages per file
func splitPDF(filePath string, pages int) error {
	ctx, err := api.ReadContextFile(filePath)
	if err != nil {
		return err
	}

	totalPages := ctx.PageCount
	dir, file := filepath.Split(filePath)
	baseName := file[:len(file)-len(filepath.Ext(file))]

	for i := 0; i < totalPages; i += pages {
		start := i + 1
		end := i + pages
		if end > totalPages {
			end = totalPages
		}

		splitFilePath := filepath.Join(dir, fmt.Sprintf("%s_%d-%d.pdf", baseName, start, end))
		err := api.ExtractPagesFile(filePath, dir, []string{fmt.Sprintf("%d-%d", start, end)}, nil)
		if err != nil {
			return err
		}

		log.Println("Created split file:", splitFilePath)
	}

	return nil
}

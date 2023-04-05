package informationPage

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func InformationPage(Window fyne.Window) {
	var productNumbersAndStatus map[string]string

	var listItems []fyne.CanvasObject

	for i, s := range productNumbersAndStatus {
		numberLabel := widget.NewLabel(i)
		statusLabel := widget.NewLabel(s)

		listItems = append(listItems, container.New(
			layout.NewHBoxLayout(),
			numberLabel,
			statusLabel,
		))
	}

	scrollContainer := container.NewVScroll(container.NewVBox(listItems...))

	Window.SetContent(scrollContainer)
}

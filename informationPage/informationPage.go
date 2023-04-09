package informationPage

import (
	. "TwistAndWrapP/pkg"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var ListInformationPage []*InformationPage

func NewInformationPage(w fyne.Window) Page {
	page := InformationPage{WindowPage: w}
	page.SetWindowContent()
	ListInformationPage = append(ListInformationPage, &page)
	w.Show()
	return page
}

type InformationPage struct {
	WindowPage fyne.Window
}

func (i InformationPage) Window() fyne.Window {
	return i.WindowPage
}

func (i InformationPage) SetWindowContent() {
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

	i.WindowPage.SetContent(scrollContainer)
}

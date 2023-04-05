package checkoutPage

import (
	. "TwistAndWrapP/pkg"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CheckoutPage(Window fyne.Window) {
	products := GetProducts()
	productCheckboxes := make([]*widget.Check, len(products))
	for i, p := range products {
		productCheckboxes[i] = widget.NewCheck(p.Name, nil)
	}

	//calculateButton := widget.NewButton("Calculate", func() {
	//	total := 0.0
	//	for i, p := range products {
	//		if productCheckboxes[i].Checked {
	//			id, _ := strconv.ParseFloat(p.Id, 64)
	//			total += id
	//		}
	//	}
	//	// Повідомляємо користувача про вартість вибраних продуктів
	//	widget.NewLabel(fmt.Sprintf("Selected products cost $%.2f", total))
	//})

	// Створюємо контейнер зі списком вибираємих продуктів та кнопкою
	var items []fyne.CanvasObject
	for _, c := range productCheckboxes {
		items = append(items, c)
	}

	content := container.NewVBox(items...)
	//content.Add(calculateButton)

	Window.SetContent(content)
}

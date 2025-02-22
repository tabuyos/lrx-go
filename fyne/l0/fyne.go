package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	counter := 0
	app := app.New()
	win := app.NewWindow("Hello tabuyos")
	win.Resize(fyne.NewSize(300, 150))
	label := widget.NewLabelWithStyle("Hello Fyne", fyne.TextAlignCenter, fyne.TextStyle{})
	quit := widget.NewButton("Quit", func() {
		app.Quit()
	})
	welcome := widget.NewButton("Welcome", func() {
		counter++
		label.SetText(fmt.Sprintf("Hello tabuyos[%d]", counter))
	})
	win.SetContent(container.NewVBox(label, welcome, quit))
	win.ShowAndRun()
}

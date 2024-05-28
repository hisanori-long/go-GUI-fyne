package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		widget.NewScrollContainer(
			widget.NewVBox(
				widget.NewButton("One", nil),
				widget.NewButton("Two", nil),
				widget.NewButton("Three", nil),
				widget.NewButton("Four", nil),
				widget.NewButton("Five", nil),
				widget.NewButton("Six", nil),
				widget.NewButton("Seven", nil),
				widget.NewButton("Eight", nil),
				widget.NewButton("Nine", nil),
				widget.NewButton("Ten", nil),
			),
		),
	)
	w.ShowAndRun()
}

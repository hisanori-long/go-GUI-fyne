package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayout(3),
			widget.NewButton("One", nil),
			widget.NewButton("Two", nil),
			widget.NewButton("Three", nil),
			widget.NewButton("Four", nil),
			layout.NewSpacer(),
			widget.NewButton("Five", nil),
			widget.NewButton("Six", nil),
			widget.NewButton("Seven", nil),
			widget.NewButton("Eight", nil),
			layout.NewSpacer(),
			widget.NewButton("Nine", nil),
			widget.NewButton("Ten", nil),
			widget.NewButton("Eleven", nil),
		),
	)
	w.ShowAndRun()
}

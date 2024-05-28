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
	bt := widget.NewButton("Top", nil)
	bb := widget.NewButton("Bottom", nil)
	bl := widget.NewButton("Left", nil)
	br := widget.NewButton("Right", nil)
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewBorderLayout(
				bt, br, bl, bb,
			),
			bt, bb, bl, br,
			widget.NewLabel("Center."),
		),
	)
	w.ShowAndRun()
}

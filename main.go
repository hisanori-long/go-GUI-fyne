package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	v := 0.
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	p := widget.NewProgressBar()
	b := widget.NewButton("Up!", func() {
		v += 0.2
		if v > 1.0 {
			v = 0.
		}
		p.SetValue(v)
	})
	w.SetContent(
		widget.NewVBox(
			l, p, b,
		),
	)

	w.ShowAndRun()
}

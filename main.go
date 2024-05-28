package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	l := widget.NewLabel("Hello Fyne!")
	ck1 := widget.NewCheck("check 1", nil)
	ck2 := widget.NewCheck("check 2", nil)

	w.SetContent(
		widget.NewVBox(
			l,
			widget.NewGroup("Group",
				ck1,
				ck2,
			),
			widget.NewButton("OK", func() {
				re := "result: "
				if ck1.Checked {
					re += "check-1 "
				}
				if ck2.Checked {
					re += "check-2 "
				}
				l.SetText(re)
			}),
		),
	)

	w.ShowAndRun()
}

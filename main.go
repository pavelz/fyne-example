package main

import "fmt"
import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	fmt.Println("vim-go")

	app := app.New()
	w := app.NewWindow("Hi")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hi macOS"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()

}

package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()

	windowConfigs(a)
}

// INFO -> Just a test function
func docsLayout(a fyne.App) {
	myWindow := a.NewWindow("Box Layout")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	myWindow.ShowAndRun()

}

func windowConfigs(a fyne.App) {

	w := a.NewWindow("Hello")

	entry := widget.NewEntry()

	hello := container.NewVBox(
		entry,
	)
	hello.Resize(fyne.NewSize(200, 100))

	status := canvas.NewText("404 not found", color.White)
	duration := canvas.NewText("0ms", color.White)
	size := canvas.NewText("0 B", color.White)
	header := container.NewHBox(
		status,
		duration,
		size,
	)
	response := container.NewVBox(
		header,
	)
	response.Resize(fyne.NewSize(300, 300))
	w.Resize(fyne.NewSize(600, 300))

	firstContainer := container.New(layout.NewHBoxLayout(), hello, layout.NewSpacer(), response)
	w.SetContent(firstContainer)
	w.ShowAndRun()

}

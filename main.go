package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/expr-lang/expr"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(200, 300))

	hello := widget.NewLabel("Calculator")
	helloCenter := container.NewCenter(hello)

	display := widget.NewLabel("######")
	displayCenter := container.NewCenter(display)

	result := widget.NewLabel("Result: 0")
	resultCenter := container.NewCenter(result)

	vbox := container.NewVBox()
	vbox.Add(helloCenter)
	vbox.Add(displayCenter)
	vbox.Add(resultCenter)

	expression := ""

	buttons := []*widget.Button{
		widget.NewButton("7", func() {
			expression += "7"
			display.SetText(expression)
		}),
		widget.NewButton("8", func() {
			expression += "8"
			display.SetText(expression)
		}),
		widget.NewButton("9", func() {
			expression += "9"
			display.SetText(expression)
		}),
		widget.NewButton("+", func() {
			expression += "+"
			display.SetText(expression)
		}),
		widget.NewButton("4", func() {
			expression += "4"
			display.SetText(expression)
		}),
		widget.NewButton("5", func() {
			expression += "5"
			display.SetText(expression)
		}),
		widget.NewButton("6", func() {
			expression += "6"
			display.SetText(expression)
		}),
		widget.NewButton("-", func() {
			expression += "-"
			display.SetText(expression)
		}),
		widget.NewButton("1", func() {
			expression += "1"
			display.SetText(expression)
		}),
		widget.NewButton("2", func() {
			expression += "2"
			display.SetText(expression)
		}),
		widget.NewButton("3", func() {
			expression += "3"
			display.SetText(expression)
		}),
		widget.NewButton("*", func() {
			expression += "*"
			display.SetText(expression)
		}),
		widget.NewButton(".", func() {
			expression += "."
			display.SetText(expression)
		}),
		widget.NewButton("0", func() {
			expression += "0"
			display.SetText(expression)
		}),
		widget.NewButton("/", func() {
			expression += "/"
			display.SetText(expression)
		}),
		widget.NewButton("=", func() {
			num, err := expr.Eval(expression, nil)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			result.SetText(fmt.Sprintf("Result: %v", num))
		}),
		widget.NewButton("<---", func() {
			expression = expression[:len(expression)-1]
			display.SetText(expression)
		}),
	}
	buttonCounter := 0
	hbox := container.NewHBox()
	center := container.NewCenter(hbox)
	vbox.Add(center)
	for _, button := range buttons {
		hbox.Add(button)
		buttonCounter += 1
		if buttonCounter >= 4 {
			buttonCounter = 0
			hbox = container.NewHBox()
			center := container.NewCenter(hbox)
			vbox.Add(center)
		}
	}

	w.SetContent(vbox)

	w.ShowAndRun()
}

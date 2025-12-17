package main

// Importing things like Fyne and EXPR
import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/expr-lang/expr"
)

// Main function, the heart of this application
func main() {
	// Create a new Fyne application
	a := app.New()

	// Create a new Fyne window
	w := a.NewWindow("Hello")

	// Resize it up to 200w 300h
	w.Resize(fyne.NewSize(200, 300))

	// Fancy calculator label on top
	calculator := widget.NewLabel("Calculator")
	calculatorCenter := container.NewCenter(calculator)

	// Label that shows the expression string
	display := widget.NewLabel("######")
	displayCenter := container.NewCenter(display)

	// Label that shows the result
	result := widget.NewLabel("Result: 0")
	resultCenter := container.NewCenter(result)

	// Vbox that'll have all the GUI elements
	vbox := container.NewVBox()
	vbox.Add(calculatorCenter)
	vbox.Add(displayCenter)
	vbox.Add(resultCenter)

	// The expression string! It'll be turned into numbers later anyway
	expression := ""

	// A huge array with buttons, includes numbers and operators
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

	// Well tricky stuff will go here.
	// I decided to write the button array in a dumb way, but here I decided
	// to take my glosses on and say "I'm an academic!11"
	// 1. Declare a counter variable, all great things start with it
	buttonCounter := 0
	// 2. Horizontal box and center container
	hbox := container.NewHBox()
	center := container.NewCenter(hbox)
	vbox.Add(center)
	// 3. Go through all buttons
	for _, button := range buttons {
		hbox.Add(button)
		buttonCounter += 1
		// 4. Each 4th button create new HBox and center container
		// This way it'll allow to place buttons on a new line,
		// which will show them in a grid-like pattern
		if buttonCounter >= 4 {
			buttonCounter = 0
			hbox = container.NewHBox()
			center := container.NewCenter(hbox)
			vbox.Add(center)
		}
	}

	// Set vbox as content of our window
	w.SetContent(vbox)

	// Finally, show what we have got
	w.ShowAndRun()
}

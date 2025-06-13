package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	password, err := GeneratePassword(20)
	if err != nil {
		panic(err)
	}
	fmt.Println(password)
	a := app.New()
	w := a.NewWindow("Password Vault")
	passwordText := widget.NewLabel("")
	passwordText.SetText("Password Soon")
	generateButton := widget.NewButton("Generate", func() {
		newPassword, err := GeneratePassword(20)
		if err == nil {
			passwordText.SetText(newPassword)
		}
	})
	content := container.NewWithoutLayout(
		passwordText,
		generateButton,
	)

	// Position widgets manually
	passwordText.Move(fyne.NewPos(225, 10))    // top left
	passwordText.Resize(fyne.NewSize(580, 30)) // make it wide enough

	generateButton.Move(fyne.NewPos(250, 150))   // bottom left (adjust Y as needed)
	generateButton.Resize(fyne.NewSize(100, 40)) // your desired size

	w.SetContent(content)
	//generateButton.Resize(fyne.NewSize(50, 50))

	w.Resize(fyne.NewSize(600, 200))
	w.ShowAndRun()
}

package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
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

	passwordTextBound := binding.NewString()
	passwordTextBound.Set("Ready to generate password...")
	passwordTextWid := widget.NewLabelWithData(passwordTextBound)

	generateButton := widget.NewButton("Generate", func() {
		newPassword, err := GeneratePassword(20)
		if err == nil {
			passwordTextBound.Set(newPassword)
		}
	})
	copyButton := widget.NewButtonWithIcon("Copy", theme.ContentCopyIcon(), func() {
		if content, err := passwordTextBound.Get(); err == nil {
			w.Clipboard().SetContent(content)
		}
	})
	content := container.NewWithoutLayout(
		passwordTextWid,
		copyButton,
		generateButton,
	)

	// Position widgets manually
	passwordTextWid.Move(fyne.NewPos(200, 10))    // top left
	passwordTextWid.Resize(fyne.NewSize(580, 30)) // make it wide enough

	generateButton.Move(fyne.NewPos(175, 150))   // bottom left (adjust Y as needed)
	generateButton.Resize(fyne.NewSize(100, 40)) // your desired size

	copyButton.Move(fyne.NewPos(325, 150))
	copyButton.Resize(fyne.NewSize(100, 40))

	w.SetContent(content)

	w.Resize(fyne.NewSize(600, 200))
	w.ShowAndRun()
}

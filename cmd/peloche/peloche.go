package main

import (
	"peloche/internal/infra/ui"
)

// https://docs.fyne.io/extend/bundle.html
//go:generate fyne bundle -o ../../internal/infra/ui/assets/assets.go --pkg assets --prefix Resource ../../assets/embed

func main() {
	ui.NewUI().Start()
}

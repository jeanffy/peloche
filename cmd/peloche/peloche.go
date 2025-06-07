package main

import (
	"peloche/internal/infra"
)

// https://docs.fyne.io/extend/bundle.html
//go:generate fyne bundle -o ../../internal/infra/ui/assets/assets.go --pkg assets --prefix Resource ../../assets/embed

func main() {
	infra.NewPelocheApp().Start()
}

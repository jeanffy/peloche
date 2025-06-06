package main

import (
	"peloche/infra"
)

// https://docs.fyne.io/extend/bundle.html
//go:generate fyne bundle -o infra/ui/assets/assets.go --pkg assets --prefix Resource etc/resources

func main() {
	infra.NewUI().Start()
}

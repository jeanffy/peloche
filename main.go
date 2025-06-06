package main

import (
	"os"
	"peloche/domain"
	"peloche/infra"
	"peloche/infra/adapters"
	"peloche/infra/ui/events"
)

// https://docs.fyne.io/extend/bundle.html
//go:generate fyne bundle -o infra/ui/assets/assets.go --pkg assets --prefix Resource etc/resources

func main() {
	var argumentPath *string = nil
	if len(os.Args) > 1 {
		argumentPath = &os.Args[1]
	}

	log := adapters.NewLogAdapter()
	fs := adapters.NewRealFsAdapter(log)

	appData := domain.NewAppData(log, fs)
	appData.SetArgumentPath(argumentPath)

	eventBus := events.NewSimpleEventBus()

	infra.NewUI().Start(appData, eventBus)
}

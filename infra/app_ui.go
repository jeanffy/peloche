package infra

import (
	"embed"
	"peloche/domain"
	"peloche/infra/ui"
	"peloche/infra/ui/context"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/lang"
)

//go:embed translation
var translations embed.FS

type AppUI struct {
}

func NewAppUI() *AppUI {
	return &AppUI{}
}

func (x *AppUI) Start(appData *domain.AppData, eventBus events.EventBus) {
	err := lang.AddTranslationsFS(translations, "translation")
	if err != nil {
		panic(err)
	}

	app := app.New()
	windowManager := ui.NewAppUIWindowManager(app)
	appUIContext := context.NewAppUIContext(app, windowManager, appData, eventBus)
	windowManager.CreateExplorerWindow(appUIContext)
}

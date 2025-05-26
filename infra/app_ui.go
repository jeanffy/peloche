package infra

import (
	"embed"
	"fmt"
	"peloche/domain"
	"peloche/infra/ui"
	"peloche/infra/ui/events"
	"peloche/infra/ui/views"

	"fyne.io/fyne/v2"
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

	fyneApp := app.New()
	fyneWin := fyneApp.NewWindow("Peloche")

	appUIContext := ui.NewAppUIContext(fyneApp, fyneWin, appData, eventBus)
	content := views.NewExplorerView(appUIContext).UIContainer

	fyneWin.SetContent(content)
	fyneWin.Resize(fyne.NewSize(900, 600))

	fyneWin.SetCloseIntercept(func() {
		fmt.Println(fyneWin.Canvas().Size().Width)
		fmt.Println(fyneWin.Canvas().Size().Height)
		fyneWin.Close()
	})

	fyneWin.ShowAndRun()
}

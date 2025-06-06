package infra

import (
	"embed"
	"fmt"
	"peloche/domain"
	"peloche/infra/ui"
	"peloche/infra/ui/context"
	"peloche/infra/ui/events"

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

	app := app.New()

	win := app.NewWindow("Péloche")

	win.SetMaster()
	win.Resize(fyne.NewSize(900, 600))

	win.SetCloseIntercept(func() {
		fmt.Println(win.Canvas().Size().Width)
		fmt.Println(win.Canvas().Size().Height)
		win.Close()
	})

	router := ui.NewAppUIRouter(win)
	dialogs := ui.NewAppUIDialogs(app, router)
	appUIContext := context.NewUIContext(app, dialogs, router, appData, eventBus)
	router.SetAppUIContext(appUIContext)
	router.NavigateToExplorerView()

	win.ShowAndRun()
}

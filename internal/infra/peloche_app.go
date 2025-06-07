package infra

import (
	"embed"
	"log"
	"peloche/internal/domain"
	"peloche/internal/infra/adapters"
	"peloche/internal/infra/ui"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/lang"
)

//go:embed translation
var translations embed.FS

type PelocheApp struct {
}

func NewPelocheApp() *PelocheApp {
	return &PelocheApp{}
}

func (x *PelocheApp) Start() {
	err := lang.AddTranslationsFS(translations, "translation")
	if err != nil {
		panic(err)
	}

	app := app.New()

	win := app.NewWindow("Péloche")

	win.SetMaster()
	win.Resize(fyne.NewSize(900, 600))

	win.SetCloseIntercept(func() {
		log.Println(win.Canvas().Size().Width)
		log.Println(win.Canvas().Size().Height)
		win.Close()
	})

	x.initDI(app, win)

	router := di.GetBasicDI().Resolve(ui.ROUTER_PORT_TOKEN).(ui.RouterPort)
	router.NavigateToExplorerView()

	win.ShowAndRun()
}

func (x *PelocheApp) initDI(app fyne.App, win fyne.Window) {
	di := di.GetBasicDI()

	di.Provide("FyneApp", app)

	di.Provide(domain.LOG_PORT_TOKEN, adapters.NewConsoleLogAdapter)
	di.Provide(domain.FS_PORT_TOKEN, adapters.NewRealFsAdapter)

	di.Provide(ui.EVENTS_PORT_TOKEN, adapters.NewSimpleEventsAdapter)

	router := adapters.NewRouterAdapter(win)
	di.Provide(ui.ROUTER_PORT_TOKEN, router)

	di.Provide(ui.DIALOGS_PORT_TOKEN, adapters.NewDialogsAdapter)

	appData := domain.NewAppData()
	di.Provide(domain.APP_DATA_TOKEN, appData)

	di.Provide(ui.CONTEXT_TOKEN, ui.NewContext(app))
}

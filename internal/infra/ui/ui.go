package ui

import (
	"embed"
	"log"
	"peloche/internal/domain"
	"peloche/internal/infra/adapters"
	"peloche/internal/infra/ui/context"
	"peloche/internal/infra/ui/dialogs"
	"peloche/internal/infra/ui/events"
	"peloche/internal/infra/ui/routing"
	"peloche/pkg/di"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/lang"
)

//go:embed translation
var translations embed.FS

type UI struct {
}

func NewUI() *UI {
	return &UI{}
}

func (x *UI) Start() {
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

	router := di.GetBasicDI().Resolve(routing.ROUTER_TOKEN).(routing.Router)
	router.NavigateToExplorerView()

	win.ShowAndRun()
}

func (x *UI) initDI(app fyne.App, win fyne.Window) {
	di := di.GetBasicDI()

	di.Provide("FyneApp", app)

	di.Provide(domain.LOG_PORT_TOKEN, adapters.NewLogAdapter)
	di.Provide(domain.FS_PORT_TOKEN, adapters.NewRealFsAdapter)

	di.Provide(events.EVENT_BUS_TOKEN, events.NewSimpleEventBus)

	router := NewUIRouter(win)
	di.Provide(routing.ROUTER_TOKEN, router)

	di.Provide(dialogs.DIALOGS_TOKEN, NewUIDialogs)

	appData := domain.NewAppData()
	di.Provide(domain.APP_DATA_TOKEN, appData)

	di.Provide(context.UI_CONTEXT_TOKEN, context.NewUIContext(app))
}

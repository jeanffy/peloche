package infra

import (
	"embed"
	"fmt"
	"peloche/domain"
	"peloche/domain/ports"
	"peloche/infra/adapters"
	"peloche/infra/ui"
	"peloche/infra/ui/context"
	"peloche/infra/ui/dialogs"
	"peloche/infra/ui/events"
	"peloche/infra/ui/routing"
	"peloche/utils"

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
		fmt.Println(win.Canvas().Size().Width)
		fmt.Println(win.Canvas().Size().Height)
		win.Close()
	})

	x.initDI(app, win)

	router := utils.GetNaiveDI().Resolve(routing.ROUTER_TOKEN).(routing.Router)
	router.NavigateToExplorerView()

	win.ShowAndRun()
}

func (x *UI) initDI(app fyne.App, win fyne.Window) {
	di := utils.GetNaiveDI()

	di.Provide("FyneApp", app)
	di.Provide(ports.LOG_PORT_TOKEN, adapters.NewLogAdapter())
	di.Provide(ports.FS_PORT_TOKEN, adapters.NewRealFsAdapter())
	di.Provide(events.EVENT_BUS_TOKEN, events.NewSimpleEventBus())
	router := ui.NewUIRouter(win)
	di.Provide(routing.ROUTER_TOKEN, router)
	di.Provide(dialogs.DIALOGS_TOKEN, ui.NewUIDialogs())

	appData := domain.NewAppData()
	di.Provide(domain.APP_DATA_TOKEN, appData)

	di.Provide(context.UI_CONTEXT_TOKEN, context.NewUIContext(app))
}

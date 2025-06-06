package context

import (
	"peloche/domain"
	"peloche/domain/ports"
	"peloche/infra/ui/dialogs"
	"peloche/infra/ui/events"
	"peloche/infra/ui/routing"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type UIContext struct {
	dialogs            dialogs.Dialogs
	router             routing.Router
	eventBus           events.EventBus
	appData            *domain.AppData
	ThemeVariant       fyne.ThemeVariant
	GridSize           uint
	GridSizeMin        uint
	GridSizeMax        uint
	SelectedPhotoIndex int
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewUIContext(fyneApp fyne.App, dialogs dialogs.Dialogs, router routing.Router, appData *domain.AppData, eventBus events.EventBus) *UIContext {
	return &UIContext{
		dialogs:            dialogs,
		router:             router,
		eventBus:           eventBus,
		appData:            appData,
		ThemeVariant:       fyneApp.Settings().ThemeVariant(),
		GridSize:           200,
		GridSizeMin:        80,
		GridSizeMax:        500,
		SelectedPhotoIndex: -1,
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *UIContext) ShowMessageBox(message string) {
	x.dialogs.MessageDialog(message)
}

func (x *UIContext) ShowErrorDialog(err error) {
	x.dialogs.ErrorDialog(err)
}

func (x *UIContext) LogInfo(params ports.LogPortParams) {
	x.appData.Log.Info(params)
}

func (x *UIContext) LogError(params ports.LogPortErrorParams) {
	x.appData.Log.Error(params)
}

func (x *UIContext) SubscribeToEvent(id string, fn interface{}) {
	x.eventBus.Subscribe(id, fn)
}

func (x *UIContext) PublishEvent(id string, args ...interface{}) {
	x.eventBus.Publish(id, args...)
}

func (x *UIContext) SetGridSize(size uint) {
	x.GridSize = size
	x.eventBus.Publish(events.EventThumbnailSizeChanged, &events.EventThumbnailSizeChangedParams{
		Size: size,
	})
}

func (x *UIContext) SetSelectedPhotoIndex(index int) {
	x.SelectedPhotoIndex = index
	x.eventBus.Publish(events.EventSelectedPhotoChanged, &events.EventSelectedPhotoChangedParams{
		Index: index,
	})
}

func (x *UIContext) NavigateTo(route routing.RouteName, args ...interface{}) {
	x.router.NavigateTo(route, args...)
}

func (x *UIContext) GetCurrentWindow() fyne.Window {
	return x.router.GetCurrentWindow()
}

func (x *UIContext) GetFolderTree() *domain.FolderTree {
	return x.appData.FolderTree
}

func (x *UIContext) SetRootFolder(rootFolderPath *string) {
	x.appData.SetRootFolder(rootFolderPath)
}

func (x *UIContext) SetCurrentFolder(folderPath *string) {
	x.appData.SetCurrentFolder(folderPath)
}

func (x *UIContext) GetPhotoList() *domain.PhotoList {
	return x.appData.PhotoList
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

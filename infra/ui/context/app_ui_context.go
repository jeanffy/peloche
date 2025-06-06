package context

import (
	"peloche/domain"
	"peloche/domain/ports"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type AppUIContext struct {
	dialogs            ContextDialogs
	router             ContextRouter
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

func NewAppUIContext(fyneApp fyne.App, dialogs ContextDialogs, router ContextRouter, appData *domain.AppData, eventBus events.EventBus) *AppUIContext {
	return &AppUIContext{
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

func (x *AppUIContext) ShowMessageBox(message string) {
	x.dialogs.MessageDialog(message)
}

func (x *AppUIContext) ShowErrorDialog(err error) {
	x.dialogs.ErrorDialog(err)
}

func (x *AppUIContext) LogInfo(params ports.LogPortParams) {
	x.appData.Log.Info(params)
}

func (x *AppUIContext) LogError(params ports.LogPortErrorParams) {
	x.appData.Log.Error(params)
}

func (x *AppUIContext) SubscribeToEvent(id string, fn interface{}) {
	x.eventBus.Subscribe(id, fn)
}

func (x *AppUIContext) PublishEvent(id string, args ...interface{}) {
	x.eventBus.Publish(id, args...)
}

func (x *AppUIContext) SetGridSize(size uint) {
	x.GridSize = size
	x.eventBus.Publish(events.EventThumbnailSizeChanged, &events.EventThumbnailSizeChangedParams{
		Size: size,
	})
}

func (x *AppUIContext) SetSelectedPhotoIndex(index int) {
	x.SelectedPhotoIndex = index
	x.eventBus.Publish(events.EventSelectedPhotoChanged, &events.EventSelectedPhotoChangedParams{
		Index: index,
	})
}

func (x *AppUIContext) NavigateTo(route Route, args ...interface{}) {
	x.router.NavigateTo(route, args...)
}

func (x *AppUIContext) GetCurrentWindow() fyne.Window {
	return x.router.GetCurrentWindow()
}

func (x *AppUIContext) GetFolderTree() *domain.FolderTree {
	return x.appData.FolderTree
}

func (x *AppUIContext) SetRootFolder(rootFolderPath *string) {
	x.appData.SetRootFolder(rootFolderPath)
}

func (x *AppUIContext) SetCurrentFolder(folderPath *string) {
	x.appData.SetCurrentFolder(folderPath)
}

func (x *AppUIContext) GetPhotoList() *domain.PhotoList {
	return x.appData.PhotoList
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

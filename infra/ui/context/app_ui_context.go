package context

import (
	"peloche/domain"
	"peloche/domain/ports"
	"peloche/infra/ui/events"

	"fyne.io/fyne/v2"
)

type AppUIContext struct {
	WinManager         AppUIContextWinManager
	EventBus           events.EventBus
	AppData            *domain.AppData
	ThemeVariant       fyne.ThemeVariant
	GridSize           float64
	SelectedPhotoIndex int
}

func NewAppUIContext(fyneApp fyne.App, winManager AppUIContextWinManager, appData *domain.AppData, eventBus events.EventBus) *AppUIContext {
	return &AppUIContext{
		WinManager:         winManager,
		EventBus:           eventBus,
		AppData:            appData,
		ThemeVariant:       fyneApp.Settings().ThemeVariant(),
		GridSize:           200,
		SelectedPhotoIndex: -1,
	}
}

func (x *AppUIContext) LogInfo(params ports.LogPortParams) {
	x.AppData.Log.Info(params)
}

func (x *AppUIContext) LogError(params ports.LogPortErrorParams) {
	x.AppData.Log.Error(params)
}

func (x *AppUIContext) SetSelectedPhotoIndex(index int) {
	x.SelectedPhotoIndex = index
	x.EventBus.Publish(events.EventSelectedPhotoChanged, &events.EventSelectedPhotoChangedParams{
		Index: index,
	})
}

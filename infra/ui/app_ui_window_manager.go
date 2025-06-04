package ui

import (
	"fmt"
	"peloche/infra/ui/context"
	"peloche/infra/ui/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	sdialog "github.com/sqweek/dialog"
)

type AppUIWindowManager struct {
	FyneApp         fyne.App
	FyneWinExplorer fyne.Window
	FyneWinEditor   fyne.Window
}

func NewAppUIWindowManager(fyneApp fyne.App) *AppUIWindowManager {
	return &AppUIWindowManager{
		FyneApp:         fyneApp,
		FyneWinExplorer: nil,
		FyneWinEditor:   nil,
	}
}

func (x *AppUIWindowManager) MessageDialog(msg string) {
	sdialog.Message("%s", msg).Info()
	// FIXME: when dialog is closed, parent window does not get the focus back
}

func (x *AppUIWindowManager) CreateExplorerWindow(appUIContext *context.AppUIContext) {
	if x.FyneWinExplorer != nil {
		x.FyneWinExplorer.RequestFocus()
		return
	}

	x.FyneWinExplorer = x.FyneApp.NewWindow("Péloche")

	content := views.NewExplorerView(x.FyneWinExplorer, appUIContext).UIContainer

	x.FyneWinExplorer.SetContent(content)
	x.FyneWinExplorer.SetMaster()
	x.FyneWinExplorer.Resize(fyne.NewSize(900, 600))

	x.FyneWinExplorer.SetCloseIntercept(func() {
		fmt.Println(x.FyneWinExplorer.Canvas().Size().Width)
		fmt.Println(x.FyneWinExplorer.Canvas().Size().Height)
		x.CloseExplorerWindow()
	})

	x.FyneWinExplorer.ShowAndRun()
}

func (x *AppUIWindowManager) CloseExplorerWindow() {
	x.FyneWinExplorer.Close()
	x.FyneWinExplorer = nil
}

func (x *AppUIWindowManager) CreateEditorWindow(appUIContext *context.AppUIContext) {
	if x.FyneWinEditor != nil {
		x.FyneWinEditor.RequestFocus()
		return
	}

	x.FyneWinEditor = x.FyneApp.NewWindow("Péloche - editor")

	content := views.NewEditorView(x.FyneWinEditor, appUIContext).UIContainer
	x.FyneWinEditor.SetContent(content)
	x.FyneWinEditor.Resize(fyne.NewSize(900, 600))

	x.FyneWinEditor.SetCloseIntercept(func() {
		x.CloseEditorWindow()
	})

	x.FyneWinEditor.Show()
}

func (x *AppUIWindowManager) CloseEditorWindow() {
	x.FyneWinEditor.Close()
	x.FyneWinEditor = nil
}

func (x *AppUIWindowManager) ErrorDialog(err error) {
	// TODO: which parent window?
	dialog.NewError(err, x.FyneWinExplorer).Show()
}

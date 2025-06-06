package ui

import (
	"peloche/infra/ui/routing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	sdialog "github.com/sqweek/dialog"
)

// ---------------------------------------------------------------------------
// definition
// ---------------------------------------------------------------------------

type AppUIDialogs struct {
	fyneApp fyne.App
	router  routing.Router
}

// ---------------------------------------------------------------------------
// constructor
// ---------------------------------------------------------------------------

func NewAppUIDialogs(fyneApp fyne.App, router routing.Router) *AppUIDialogs {
	return &AppUIDialogs{
		fyneApp: fyneApp,
		router:  router,
	}
}

// ---------------------------------------------------------------------------
// public
// ---------------------------------------------------------------------------

func (x *AppUIDialogs) MessageDialog(msg string) {
	sdialog.Message("%s", msg).Info()
	// FIXME: when dialog is closed, parent window does not get the focus back
}

func (x *AppUIDialogs) ErrorDialog(err error) {
	parent := x.router.GetCurrentWindow()
	dialog.NewError(err, parent).Show()
}

// ---------------------------------------------------------------------------
// events
// ---------------------------------------------------------------------------

// ---------------------------------------------------------------------------
// private
// ---------------------------------------------------------------------------

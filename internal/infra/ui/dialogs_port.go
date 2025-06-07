package ui

var DIALOGS_PORT_TOKEN = "DialogsPort"

type DialogsPort interface {
	MessageDialog(msg string)
	ErrorDialog(err error)
}

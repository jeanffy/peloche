package dialogs

var DIALOGS_TOKEN = "Dialogs"

type Dialogs interface {
	MessageDialog(msg string)
	ErrorDialog(err error)
}

package dialogs

type Dialogs interface {
	MessageDialog(msg string)
	ErrorDialog(err error)
}

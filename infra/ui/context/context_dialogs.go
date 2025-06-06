package context

type ContextDialogs interface {
	MessageDialog(msg string)
	ErrorDialog(err error)
}

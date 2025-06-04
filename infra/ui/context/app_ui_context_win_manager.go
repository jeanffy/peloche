package context

type AppUIContextWinManager interface {
	MessageDialog(msg string)
	ErrorDialog(err error)

	CreateExplorerWindow(appUIContext *AppUIContext)
	CloseExplorerWindow()
	CreateEditorWindow(appUIContext *AppUIContext)
	CloseEditorWindow()
}

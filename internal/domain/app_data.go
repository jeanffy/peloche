package domain

var APP_DATA_TOKEN = "AppData"

// ---------------------------------------------------------------------------
// #region definition

type AppData struct {
	ArgumentPath      *string
	RootFolderPath    *string
	FolderTree        *FolderTree
	CurrentFolderPath *string
	PhotoList         *PhotoList
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewAppData() *AppData {
	return &AppData{}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *AppData) SetArgumentPath(argumentPath *string) {
	x.ArgumentPath = argumentPath
}

func (x *AppData) SetRootFolder(rootFolderPath *string) {
	x.RootFolderPath = rootFolderPath
	x.CurrentFolderPath = x.RootFolderPath
	x.FolderTree = NewFolderTree(*x.RootFolderPath, 0)
}

func (x *AppData) SetCurrentFolder(currentFolderPath *string) {
	x.CurrentFolderPath = currentFolderPath
	x.PhotoList = NewPhotoList(*x.CurrentFolderPath)
}

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion

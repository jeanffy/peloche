package domain

import (
	"peloche/domain/ports"
)

type AppData struct {
	Log ports.LogPort
	Fs  ports.FsPort

	ArgumentPath      *string
	RootFolderPath    *string
	FolderTree        *FolderTree
	CurrentFolderPath *string
	PhotoList         *PhotoList
}

func NewAppCore(log ports.LogPort, fs ports.FsPort) *AppData {
	return &AppData{
		Log: log,
		Fs:  fs,
	}
}

func (x *AppData) SetArgumentPath(argumentPath *string) {
	x.ArgumentPath = argumentPath
}

func (x *AppData) SetRootFolder(rootFolderPath *string) {
	x.RootFolderPath = rootFolderPath
	x.CurrentFolderPath = x.RootFolderPath
	x.FolderTree = NewFolderTree(x.Log, x.Fs, *x.RootFolderPath, 0)
}

func (x *AppData) SetCurrentFolder(currentFolderPath *string) {
	x.CurrentFolderPath = currentFolderPath
	x.PhotoList = NewPhotoList(x.Log, x.Fs, *x.CurrentFolderPath)
}

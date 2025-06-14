package domain

import (
	"path/filepath"
	"peloche/pkg/di"
	"reflect"
)

// ---------------------------------------------------------------------------
// #region definition

type FolderTree struct {
	Name       string
	Path       string
	SubFolders []*FolderTree
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewFolderTree(folderPath string, level int) *FolderTree {
	if level > 5 {
		return &FolderTree{
			Name:       filepath.Base(folderPath),
			Path:       folderPath,
			SubFolders: []*FolderTree{},
		}
	}

	log := di.GetBasicDI().Resolve(LOG_PORT_TOKEN).(LogPort)
	fs := di.GetBasicDI().Resolve(FS_PORT_TOKEN).(FsPort)

	entries, err := fs.ReadDir(folderPath)
	if err != nil {
		log.Error(LogPortErrorParams{
			Module: reflect.TypeOf(FolderTree{}).String(),
			Error:  err,
			Msg:    folderPath,
		})
		return &FolderTree{
			Name:       filepath.Base(folderPath),
			Path:       folderPath,
			SubFolders: []*FolderTree{},
		}
	}

	subFolders := make([]*FolderTree, 0, len(entries))
	for _, e := range entries {
		if e.IsDir {
			subFolderPath := filepath.Join(folderPath, e.Name)
			subFolderTree := NewFolderTree(subFolderPath, level+1)
			subFolders = append(subFolders, subFolderTree)
		}
	}

	return &FolderTree{
		Name:       filepath.Base(folderPath),
		Path:       folderPath,
		SubFolders: subFolders,
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *FolderTree) Find(path string) *FolderTree {
	if x.Path == path {
		return x
	}
	for _, item := range x.SubFolders {
		found := item.Find(path)
		if found != nil {
			return found
		}
	}
	return nil
}

func (x *FolderTree) GetSubFolderPaths() []string {
	paths := make([]string, 0, len(x.SubFolders))
	for _, item := range x.SubFolders {
		paths = append(paths, item.Path)
	}
	return paths
}

func (x *FolderTree) HasSubFolders() bool {
	return len(x.SubFolders) > 0
}

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion

package domain

import (
	"path/filepath"
	"peloche/pkg/di"
	"reflect"
	"slices"
	"strings"
)

var HANDLED_EXTENSIONS = []string{".heic", ".jpg", ".jpeg"}

// ---------------------------------------------------------------------------
// #region definition

type PhotoList struct {
	Photos []*Photo
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func NewPhotoList(folderPath string) *PhotoList {
	log := di.GetBasicDI().Resolve(LOG_PORT_TOKEN).(LogPort)
	fs := di.GetBasicDI().Resolve(FS_PORT_TOKEN).(FsPort)

	entries, err := fs.ReadDir(folderPath)
	if err != nil {
		log.Error(LogPortErrorParams{
			Module: reflect.TypeOf(PhotoList{}).Name(),
			Error:  err,
			Msg:    folderPath,
		})
		return &PhotoList{
			Photos: []*Photo{},
		}
	}

	photos := make([]*Photo, 0, len(entries))
	for _, e := range entries {
		if e.IsFile && slices.Contains(HANDLED_EXTENSIONS, strings.ToLower(e.Ext)) {
			filePath := filepath.Join(folderPath, e.Name)
			photos = append(photos, NewPhoto(e.Name, e.Ext, filePath))
		}
	}

	return &PhotoList{
		Photos: photos,
	}
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion

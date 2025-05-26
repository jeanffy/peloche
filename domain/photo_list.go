package domain

import (
	"path/filepath"
	"peloche/domain/ports"
	"reflect"
	"slices"
	"strings"
)

var handledExtensions = []string{".heic", ".jpg", ".jpeg"}

type PhotoList struct {
	log ports.LogPort
	fs  ports.FsPort

	Photos []*Photo
}

func NewPhotoList(log ports.LogPort, fs ports.FsPort, folderPath string) *PhotoList {
	entries, err := fs.ReadDir(folderPath)
	if err != nil {
		log.Error(ports.LogPortErrorParams{
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
		if e.IsFile && slices.Contains(handledExtensions, strings.ToLower(e.Ext)) {
			filePath := filepath.Join(folderPath, e.Name)
			photos = append(photos, NewPhoto(e.Name, e.Ext, filePath))
		}
	}

	return &PhotoList{
		log:    log,
		fs:     fs,
		Photos: photos,
	}
}

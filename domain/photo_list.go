package domain

import (
	"path/filepath"
	"peloche/domain/ports"
	"peloche/utils"
	"reflect"
	"slices"
	"strings"
)

var handledExtensions = []string{".heic", ".jpg", ".jpeg"}

type PhotoList struct {
	Photos []*Photo
}

func NewPhotoList(folderPath string) *PhotoList {
	log := utils.GetNaiveDI().Resolve(ports.LOG_PORT_TOKEN).(ports.LogPort)
	fs := utils.GetNaiveDI().Resolve(ports.FS_PORT_TOKEN).(ports.FsPort)

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
		Photos: photos,
	}
}

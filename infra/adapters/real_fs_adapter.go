package adapters

import (
	"os"
	"path/filepath"
	"peloche/domain/ports"
	"peloche/internal/di"
	"reflect"
)

type RealFsAdapter struct {
	log ports.LogPort
}

func NewRealFsAdapter() *RealFsAdapter {
	return &RealFsAdapter{
		log: di.GetBasicDI().Resolve(ports.LOG_PORT_TOKEN).(ports.LogPort),
	}
}

func (x *RealFsAdapter) ReadDir(dirPath string) ([]ports.FsPortEntry, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		x.log.Error(ports.LogPortErrorParams{
			Module: reflect.TypeOf(RealFsAdapter{}).String(),
			Error:  err,
			Msg:    dirPath,
		})
		return []ports.FsPortEntry{}, nil
	}

	mapped := make([]ports.FsPortEntry, 0, len(entries))
	for _, entry := range entries {
		mapped = append(mapped, ports.FsPortEntry{
			IsDir:      entry.IsDir(),
			IsFile:     !entry.IsDir(),
			Name:       entry.Name(),
			Ext:        filepath.Ext(entry.Name()),
			ParentPath: dirPath,
		})
	}
	return mapped, nil
}

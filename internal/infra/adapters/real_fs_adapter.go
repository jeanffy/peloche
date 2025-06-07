package adapters

import (
	"os"
	"path/filepath"
	"peloche/internal/domain"
	"peloche/pkg/di"
	"reflect"
)

type RealFsAdapter struct {
	log domain.LogPort
}

func NewRealFsAdapter() *RealFsAdapter {
	return &RealFsAdapter{
		log: di.GetBasicDI().Resolve(domain.LOG_PORT_TOKEN).(domain.LogPort),
	}
}

func (x *RealFsAdapter) ReadDir(dirPath string) ([]domain.FsPortEntry, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		x.log.Error(domain.LogPortErrorParams{
			Module: reflect.TypeOf(RealFsAdapter{}).String(),
			Error:  err,
			Msg:    dirPath,
		})
		return []domain.FsPortEntry{}, nil
	}

	mapped := make([]domain.FsPortEntry, 0, len(entries))
	for _, entry := range entries {
		mapped = append(mapped, domain.FsPortEntry{
			IsDir:      entry.IsDir(),
			IsFile:     !entry.IsDir(),
			Name:       entry.Name(),
			Ext:        filepath.Ext(entry.Name()),
			ParentPath: dirPath,
		})
	}
	return mapped, nil
}

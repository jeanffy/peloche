package ports

var FS_PORT_TOKEN = "FsPort"

type FsPortEntry struct {
	IsDir      bool
	IsFile     bool
	Name       string
	Ext        string
	ParentPath string
}

type FsPort interface {
	ReadDir(dirPath string) ([]FsPortEntry, error)
}

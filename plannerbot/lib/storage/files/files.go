package files

import (
	"os"
	"path/filepath"
	e "planner/lib/e"
	"planner/lib/storage"
)

type Storage struct {
	basePath string
}

const(
	defaultPermition = 0774
)

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = e.WrapIFError("can't do request", err) }()

	filePath := filepath.Join(s.basePath, page.UserName)

	if err:=os.MkdirAll(path, defaultPermition)
}

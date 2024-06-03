package utils

import (
	"os"
	"path/filepath"
)

type FileStore struct {
	workdir string
	logger  Logger
}

func NewFileStore(workdir string, logger Logger) (*FileStore, error) {
	if workdir == "" {
		var err error
		workdir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}

	return &FileStore{
		logger:  logger,
		workdir: workdir,
	}, nil
}

func (f *FileStore) GetWorkdir() string {
	return f.workdir
}

func (f *FileStore) SetWorkdir(newWd string) {
	f.workdir = newWd
}

func (f *FileStore) Canonical(path string) string {
	path = filepath.Clean(path)

	resolvedPath, err := filepath.EvalSymlinks(path)
	if err != nil {
		f.logger.Err(err)
	} else {
		path = resolvedPath
	}

	return path
}

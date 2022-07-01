package service

import (
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
)

type File struct {
	dFile   *dao.File
	dMatter *dao.Matter
}

func NewFile() *File {
	return &File{
		dFile: dao.NewFile(),
	}
}

func (s *File) Create(uid int64, signature string, size int64) (file *model.File, err error) {
	file = model.NewFile(uid, signature, size)
	err = s.dFile.Create(file)
	return
}

func (s *File) Exist(signature string, size int64) (file *model.File, exist bool) {
	file, exist = s.dFile.Exist(signature, size)
	return file, exist
}

package dao

import (
	"errors"
	"fmt"
	"github.com/amingze/gochive/internal/app/model"
	"gorm.io/gorm"
)

type File struct {
}

func NewFile() *File {
	return &File{}
}
func (ms *File) Exist(signature string, size int64) (*model.File, bool) {
	m := new(model.File)
	err := gdb.Where("signature=? and size=?", signature, size).First(m).Error
	return m, !errors.Is(err, gorm.ErrRecordNotFound)
}

func (ms *File) Create(f *model.File) error {
	if _, ok := ms.Exist(f.Signature, f.Size); ok {
		return fmt.Errorf("file already exist")
	}
	return gdb.Create(f).Error
}

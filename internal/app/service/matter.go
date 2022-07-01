package service

import (
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
)

type Matter struct {
	dMatter *dao.Matter
}

func NewMatter() *Matter {
	return &Matter{
		dMatter: dao.NewMatter(),
	}
}
func (s *Matter) Create(m *model.Matter) (err error) {
	err = s.dMatter.Create(m)
	return
}

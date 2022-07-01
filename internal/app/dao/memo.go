package dao

import (
	"github.com/amingze/gochive/internal/app/model"
)

type Memo struct {
}

func NewMemo() *Memo {
	return &Memo{}
}

func (s *Memo) Create(memo *model.Memo) error {
	return gdb.Create(memo).Error
}

func (s *Memo) UpdateContent(id int64, content string) error {

	m, err := s.Find(id)
	if err != nil {
		return err
	}
	if err := gdb.Model(m).Update("content", content).Error; err != nil {
		return err
	}

	return nil
}
func (s *Memo) Delete(id int64) error {
	share := new(model.Memo)
	return gdb.Delete(share, id).Error
}
func (m *Memo) FindAll(q *Query) (list []model.Memo, total int64, err error) {
	sn := gdb.Where(q.SQL(), q.Params...)
	sn.Model(model.Memo{}).Count(&total)
	sn = sn.Order("created_at desc")
	if q.Offset > 0 {
		sn = sn.Offset(q.Offset)
	}
	if q.Limit > 0 {
		sn = sn.Limit(q.Limit)
	}
	err = sn.Find(&list).Error
	return
}

func (s *Memo) Find(id int64) (memo *model.Memo, err error) {
	memo = new(model.Memo)
	err = gdb.First(memo, id).Error
	return
}

func (s *Memo) FindByAlias(alias string) (memo *model.Memo, err error) {
	memo = new(model.Memo)
	err = gdb.First(memo, "alias=?", alias).Error
	return
}

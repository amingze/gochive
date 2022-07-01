package service

import (
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
)

type Memo struct {
	dMemo *dao.Memo
}

func NewMemo() *Memo {
	return &Memo{
		dMemo: dao.NewMemo(),
	}
}

func (m *Memo) Create(memo *model.Memo) error {
	return m.dMemo.Create(memo)
}

func (m *Memo) Find(id int64) (*model.Memo, error) {
	return m.dMemo.Find(id)
}

func (m *Memo) Search(uid int64, content string, offset int, limit int) (list []model.Memo, total int64, err error) {
	query := dao.NewQuery()
	query.WithEq("uid", uid)
	query.WithLike("content", content)
	query.Offset = offset
	query.Limit = limit
	return m.dMemo.FindAll(query)
}

func (m *Memo) FindAll(uid int64, offset int, limit int) (list []model.Memo, total int64, err error) {
	query := dao.NewQuery()
	query.WithEq("uid", uid)
	query.Offset = offset
	query.Limit = limit
	return m.dMemo.FindAll(query)
}

func (m *Memo) Delete(id int64) error {
	return m.dMemo.Delete(id)
}

func (m *Memo) Update(id int64, model *model.Memo) error {
	return m.dMemo.UpdateContent(id, model.Content)

}

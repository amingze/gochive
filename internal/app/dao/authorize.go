package dao

import (
	"errors"
	"fmt"
	"github.com/amingze/gochive/internal/app/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Authorize struct {
}

func NewAuthorize() *Authorize {
	return &Authorize{}
}

func (u *Authorize) Find(uid int64, name string) (*model.Authorize, error) {
	uk := new(model.Authorize)
	if err := gdb.Where("uid=? and name=?", uid, name).First(uk).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("access key not exist")
	}

	return uk, nil
}

func (u *Authorize) FindByClientID(clientID string) (*model.Authorize, error) {
	uk := new(model.Authorize)
	if err := gdb.Where("access_key=?", clientID).First(uk).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("access key not exist")
	}

	return uk, nil
}

func (u *Authorize) FindAll(query *Query) (list []*model.Authorize, total int64, err error) {
	sn := gdb.Model(&model.Authorize{})
	if len(query.Params) > 0 {
		sn = sn.Where(query.SQL(), query.Params...)
	}
	sn.Count(&total)
	err = sn.Offset(query.Offset).Limit(query.Limit).Preload(clause.Associations).Find(&list).Error
	return
}

func (u *Authorize) Create(uk *model.Authorize) (*model.Authorize, error) {
	if _, err := u.Find(uk.Uid, uk.Name); err == nil {
		return nil, fmt.Errorf("userKey already exist: %s", uk.Name)
	}

	if err := gdb.Create(uk).Error; err != nil {
		return nil, err
	}

	return uk, nil
}

func (u *Authorize) Update(user *model.Authorize) error {
	return gdb.Save(user).Error
}

func (u *Authorize) Delete(user *model.Authorize) error {
	return gdb.Delete(user).Error
}

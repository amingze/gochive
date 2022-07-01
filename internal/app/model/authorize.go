package model

import (
	"fmt"
	"github.com/amingze/gochive/internal/pkg/utils/strutil"
	"strings"
	"time"
)

type Authorize struct {
	Model
	Uid       int64  `json:"uid" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	AccessKey string `json:"access_key" gorm:"size:32;not null"`
	SecretKey string `json:"secret_key" gorm:"size:64;not null"`
}

func NewAuthorize(uid int64, name string) *Authorize {
	uk := &Authorize{
		Uid:       uid,
		Name:      name,
		AccessKey: strutil.Md5Hex(fmt.Sprintf("%d:%d:%s", uid, time.Now().Unix(), strutil.RandomText(5))),
	}
	uk.ResetSecret()
	return uk
}

func (Authorize) TableName() string {
	return "authorize"
}

func (uk *Authorize) ResetSecret() {
	l := strutil.Md5HexShort(strutil.RandomText(8))
	r := strutil.Md5HexShort(strutil.RandomText(8))
	m := strutil.Md5HexShort(l + uk.AccessKey + r)
	uk.SecretKey = strings.ToLower(l + m + r)
}

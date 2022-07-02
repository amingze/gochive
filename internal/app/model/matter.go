package model

import (
	"github.com/amingze/gochive/internal/pkg/utils/strutil"
	"path/filepath"
	"strings"
	"time"
)

var DocTypes = []string{
	"text/csv",
	"application/msword",
	"application/vnd.ms-excel",
	"application/vnd.ms-powerpoint",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation",
}

const (
	DirTypeSys = iota + 1
	DirTypeUser
	DirFileMaxNum = 65534
)

type Matter struct {
	Model
	Uid        int64      `json:"uid" gorm:"not null"`
	Fid        int64      `json:"fid" gorm:"not null"`
	Alias      string     `json:"alias" gorm:"size:16;not null"`
	Name       string     `json:"name" gorm:"not null"`
	Type       string     `json:"type" gorm:"not null"`
	DirType    int8       `json:"dirtype" gorm:"column:dirtype;not null"`
	Parent     string     `json:"parent" gorm:"not null"`
	Object     string     `json:"object" gorm:"not null"`
	URL        string     `json:"url" gorm:"-"`
	UploadedAt *time.Time `json:"uploaded"`
	TrashedBy  string     `json:"-" gorm:"size:16;not null"`
	IsFast     bool       `json:"is_fast" gorm:"-"`
}

func NewMatter(uid int64, name string) *Matter {
	return &Matter{
		Uid:  uid,
		Name: strings.TrimSpace(name),
	}
}
func (Matter) TableName() string {
	return "matter"
}

func (m *Matter) Clone() *Matter {
	clone := *m
	clone.ID = 0
	clone.Alias = strutil.RandomText(16)
	return &clone
}

func (m *Matter) FullPath() string {
	fp := m.Parent + m.Name
	if m.IsDir() {
		fp += "/"
	}
	return fp
}

func (m *Matter) IsDir() bool {
	return m.DirType > 0
}

func (m *Matter) UserAccessible(uid int64) bool {
	return m.Uid == uid
}

func (m *Matter) BuildObject(rootPath string, filePath string) {
	if filePath == "" {
		filePath = "$NOW_DATE/$RAND_16KEY.$RAW_EXT"
	}
	m.Object = filepath.Join(rootPath, m.renderPath(filePath))
}

func (m *Matter) renderPath(path string) string {
	ons := make([]string, 0)
	for _, env := range SupportEnvs {
		ons = append(ons, env.Name, env.buildV(m))
	}
	return strings.NewReplacer(ons...).Replace(path)
}

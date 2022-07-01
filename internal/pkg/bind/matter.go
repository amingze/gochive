package bind

import (
	"github.com/amingze/gochive/internal/app/model"
	"mime"
	"path/filepath"
)

type QueryFiles struct {
	QueryPage
	Sid     int64  `form:"sid" binding:"required"`
	Dir     string `form:"dir"`
	Type    string `form:"type"`
	Keyword string `form:"kw"`
}

type BodyMatter struct {
	Fid       int64  `json:"fid"`
	Name      string `json:"name"`
	IsDir     bool   `json:"is_dir"`
	Dir       string `json:"dir"`
	Type      string `json:"type"`
	Size      int64  `json:"size"`
	Signature string `json:"signature"`
}

func Body2Matter(m *BodyMatter, uid int64) (matter *model.Matter) {
	matter = model.NewMatter(uid, m.Name)
	matter.Type = m.Type
	matter.Name = m.Name
	matter.Fid = m.Fid
	return
}

func (p *BodyMatter) ToMatter(uid int64) *model.Matter {
	detectType := func(name string) string {
		cType := mime.TypeByExtension(filepath.Ext(p.Name))
		if cType != "" {
			return cType
		}

		return "application/octet-stream"
	}

	m := model.NewMatter(uid, p.Name)
	m.Type = p.Type
	//m.Size = p.Size
	m.Parent = p.Dir
	if p.IsDir {
		m.DirType = model.DirTypeUser
	} else if p.Type == "" {
		m.Type = detectType(p.Name)
	}

	return m
}

type BodyFileRename struct {
	NewName string `json:"name" binding:"required"`
}

type BodyFileMove struct {
	NewDir string `json:"dir"`
}

type BodyFileCopy struct {
	NewPath string `json:"path" binding:"required"`
}

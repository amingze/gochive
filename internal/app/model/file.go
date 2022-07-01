package model

type File struct {
	Model
	FirstAuthor   int64  `json:"signature" gorm:"not null"`
	Signature     string `json:"signature" gorm:"not null"`
	SignatureType string `json:"signature" gorm:"not null;default:'md5'"`
	Size          int64  `json:"size" gorm:"not null;default:0"`
}

func NewFile(uid int64, signature string, size int64) *File {
	return &File{
		FirstAuthor: uid,
		Signature:   signature,
		Size:        size,
	}
}

func (File) TableName() string {
	return "file"
}

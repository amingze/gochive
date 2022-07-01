package model

type UserProfile struct {
	Model
	Uid      int64  `json:"uid" gorm:"unique_index;not null"`
	Nickname string `json:"nickname" gorm:"size:32;not null"`
	Avatar   string `json:"avatar" gorm:"size:255;not null"`
	Bio      string `json:"bio" gorm:"size:255;not null"`
	URL      string `json:"url" gorm:"size:255;not null"`
	Company  string `json:"company" gorm:"size:32;not null"`
	Location string `json:"location" gorm:"size:32;not null"`
	Locale   string `json:"locale" gorm:"not null"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}

package model

import (
	"strconv"
	"strings"
)

const (
	RoleAdmin  = "admin"
	RoleMember = "member"
	RoleGuest  = "guest"
)

var roles = map[string]string{
	RoleAdmin:  "管理员",
	RoleMember: "注册用户",
	RoleGuest:  "游客",
}

type UserCreateOption struct {
	Roles      string
	Ticket     string
	Origin     string
	StorageMax uint64
}

func NewUserCreateOption() UserCreateOption {
	return UserCreateOption{}
}

type User struct {
	Model
	Email    string      `json:"email" gorm:"size:32;unique_index;not null"`
	Username string      `json:"username" gorm:"size:20;unique_index;not null"`
	Password string      `json:"-" gorm:"size:32;not null"`
	Roles    string      `json:"-" gorm:"size:64;not null"`
	RoleTxt  string      `json:"role" gorm:"-"`
	Ticket   string      `json:"ticket" gorm:"size:6;unique_index;not null"`
	Profile  UserProfile `json:"profile,omitempty" gorm:"foreignKey:Uid"`
	Token    string      `json:"-" gorm:"-"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) IDString() string {
	return strconv.FormatInt(u.ID, 10)
}

func (u *User) RolesSplit() []string {
	return strings.Split(u.Roles, ",")
}

func (u *User) Format() *User {
	u.RoleTxt = roles[u.Roles]
	return u
}

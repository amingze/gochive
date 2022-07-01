package service

import (
	"fmt"
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/pkg/utils/regexputil"
	"github.com/amingze/gochive/internal/pkg/utils/strutil"
)

type User struct {
	dUser *dao.User
	dOpt  *dao.Option

	sToken *Token
}

func NewUser() *User {
	return &User{
		dUser:  dao.NewUser(),
		dOpt:   dao.NewOption(),
		sToken: NewToken(),
	}
}

func (u *User) Signup(email, password string, opt model.UserCreateOption) (*model.User, error) {
	if _, exist := u.dUser.TicketExist(opt.Ticket); !exist && opt.Ticket != "" {
		return nil, fmt.Errorf("invalid ticket")
	}

	user := &model.User{
		Email:    email,
		Username: fmt.Sprintf("mu%s", strutil.RandomText(18)),
		Password: strutil.Md5Hex(password),
		Roles:    opt.Roles,
		Ticket:   strutil.RandomText(6),
	}
	mUser, err := u.dUser.Create(user, opt.StorageMax)
	if err != nil {
		return nil, err
	}

	return mUser, nil
}

func (u *User) SignIn(usernameOrEmail, password string, ttl int) (*model.User, error) {
	userFinder := u.dUser.UsernameExist
	if regexputil.EmailRegex.MatchString(usernameOrEmail) {
		userFinder = u.dUser.EmailExist
	}

	user, exist := userFinder(usernameOrEmail)
	if !exist {
		return nil, fmt.Errorf("用户不存在")
	} else if user.Password != strutil.Md5Hex(password) {
		return nil, fmt.Errorf("密码错误")
	}

	token, err := u.sToken.Create(user.IDString(), ttl, user.Roles)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}

func (u *User) PasswordUpdate(uid int64, oldPwd, newPwd string) error {
	user, err := u.dUser.Find(uid)
	if err != nil {
		return err
	} else if user.Password != strutil.Md5Hex(oldPwd) {
		return fmt.Errorf("error password")
	}

	user.Password = strutil.Md5Hex(newPwd)
	return u.dUser.Update(user)
}

func (u *User) PasswordResetApply(origin, email string) error {
	user, ok := u.dUser.EmailExist(email)
	if !ok {
		return fmt.Errorf("email not exist")
	}

	// issue a short-term token for password reset
	_, err := u.sToken.Create(user.IDString(), 300)
	if err != nil {
		return err
	}

	return err
}

func (u *User) PasswordReset(token, password string) error {
	rc, err := u.sToken.Verify(token)
	if err != nil {
		return err
	}

	return u.dUser.PasswordReset(rc.Uid(), password)
}

func (u *User) InviteRequired() bool {
	opts, err := u.dOpt.Get(model.OptSite)
	if err != nil {
		return false
	}

	return opts.GetBool("invite_required")
}

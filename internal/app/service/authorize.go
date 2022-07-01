package service

import (
	"context"
	"encoding/base64"
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
	"github.com/go-oauth2/oauth2/v4"
	"log"
	"strings"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

var cs = store.NewClientStore()

type Authorize struct {
	dAuthorize *dao.Authorize

	sToken *Token
}

func NewAuthorize() *Authorize {
	return &Authorize{
		dAuthorize: dao.NewAuthorize(),

		sToken: NewToken(),
	}
}

func (uk *Authorize) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	muk, err := uk.dAuthorize.FindByClientID(data.Client.GetID())
	if err != nil {
		return "", "", err
	}

	user, err := dao.NewUser().Find(muk.Uid)
	if err != nil {
		return "", "", err
	}

	ttl := data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix()
	access, err = uk.sToken.Create(user.IDString(), int(ttl), user.Roles)
	if err != nil {
		return
	}

	if isGenRefresh {
		t := uuid.NewSHA1(uuid.Must(uuid.NewRandom()), []byte(access)).String()
		refresh = base64.URLEncoding.EncodeToString([]byte(t))
		refresh = strings.ToUpper(strings.TrimRight(refresh, "="))
	}
	return
}

func (uk *Authorize) ClientStore() *store.ClientStore {
	return cs
}

func (uk *Authorize) Create(muk *model.Authorize) error {
	if _, err := uk.dAuthorize.Create(muk); err != nil {
		return err
	}

	return uk.ClientStore().Set(muk.AccessKey, &models.Client{ID: muk.AccessKey, Secret: muk.SecretKey})
}

func (uk *Authorize) ResetSecret(muk *model.Authorize) error {
	muk.ResetSecret()
	if err := uk.dAuthorize.Update(muk); err != nil {
		return err
	}

	return uk.ClientStore().Set(muk.AccessKey, &models.Client{ID: muk.AccessKey, Secret: muk.SecretKey})
}

func (uk *Authorize) LoadExistClient() {
	if !viper.IsSet("installed") {
		return
	}

	list, _, err := uk.dAuthorize.FindAll(dao.NewQuery())
	if err != nil {
		log.Println(err)
		return
	}

	for _, muk := range list {
		cli := &models.Client{ID: muk.AccessKey, Secret: muk.SecretKey}
		if err := uk.ClientStore().Set(muk.AccessKey, cli); err != nil {
			log.Println(err)
		}
	}
}

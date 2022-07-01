package v1

import (
	"fmt"
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/app/service"
	"github.com/amingze/gochive/internal/pkg/bind"
	"github.com/amingze/gochive/internal/pkg/middleware"
	"github.com/amingze/gochive/internal/pkg/utils/fileutil"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"github.com/amingze/gochive/internal/pkg/utils/jwtutil"
	"github.com/amingze/gochive/internal/pkg/utils/strutil"
	"github.com/amingze/gochive/pkg/config"

	"github.com/gin-gonic/gin"

	"github.com/spf13/viper"
)

type Option struct {
	jwtutil.JWTUtil

	sOption *service.Option
}

func NewOptionResource() *Option {
	return &Option{
		sOption: service.NewOption(),
	}
}

func (rs *Option) Register(router *gin.RouterGroup) {
	router.PUT("/system/database", rs.setupDatabase)
	router.PUT("/system/account", rs.createAdministrator)

	router.Use(middleware.Installer)
	router.Use(middleware.LoginAuth())
	router.GET("/system/matter-path-envs", rs.matterPathEnvs)
	router.GET("/system/options/:name", rs.find)
	router.PUT("/system/options/:name", rs.update)
}

// findAll godoc
// @Tags System
// @Summary 初始化数据库
// @Description 初始化数据库链接
// @Accept json
// @Produce json
// @Param query query bind.BodySystemDatabase true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /system/database [put]
func (rs *Option) setupDatabase(c *gin.Context) {
	if viper.IsSet("installed") {
		ginutil.JSONBadRequest(c, fmt.Errorf("datebase config already installed"))
		return
	}

	//p := make(map[string]string)
	p := new(bind.BodySystemDatabase)
	if err := c.ShouldBind(&p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	if err := dao.Init(p.Driver, p.DSN); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	if err := dao.NewOption().Init(); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	viper.Set("database.driver", p.Driver)
	viper.Set("database.dsn", p.DSN)
	cfgFile := viper.ConfigFileUsed()
	if cfgFile == "" {
		cfgFile = config.DefaultConfigFile
	}
	fileutil.MkFileLastDir(cfgFile)
	if err := viper.WriteConfigAs(cfgFile); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSON(c)
}

// findAll godoc
// @Tags System
// @Summary 初始化管理员账号
// @Description 初始化管理员账号
// @Accept json
// @Produce json
// @Param query query bind.BodyUserCreation true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /system/account [put]
func (rs *Option) createAdministrator(c *gin.Context) {
	if viper.IsSet("installed") {
		ginutil.JSONBadRequest(c, fmt.Errorf("datebase config already installed"))
		return
	}

	p := new(bind.BodyUserCreation)
	if err := c.ShouldBind(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}
	// 创建基本信息
	user := &model.User{
		Email:    p.Email,
		Username: "admin",
		Password: strutil.Md5Hex(p.Password),
		Roles:    "admin",
		Ticket:   strutil.RandomText(6),
	}
	if _, err := dao.NewUser().Create(user, 0); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}
	viper.Set("installed", true)

	ginutil.JSON(c)

}

// find godoc
// @Tags System
// @Summary 查找设置
// @Description 查找设置
// @Accept json
// @Produce json
// @Param name path string true "参数"
// @Security OAuth2Application[admin]
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /system/options/{name} [get]
func (rs *Option) find(c *gin.Context) {
	ret, err := dao.NewOption().Get(c.Param("name"))
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	ginutil.JSONData(c, ret)
}

// update godoc
// @Tags System
// @Summary 更新设置
// @Description 更新设置
// @Accept json
// @Produce json
// @Param name path string true "名称"
// @Param body body gin.H true "参数"
// @Security OAuth2Application[admin]
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /system/options/{name} [put]
func (rs *Option) update(c *gin.Context) {
	p := make(map[string]interface{})
	if err := c.ShouldBind(&p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	if err := rs.sOption.Update(c.Param("name"), p); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSON(c)
}

// matterPathEnvs godoc
// @Tags System
// @Summary 环境变量
// @Description 环境变量
// @Accept json
// @Produce json
// @Security OAuth2Application[admin]
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /system/matter-path-envs [get]
func (rs *Option) matterPathEnvs(c *gin.Context) {
	ginutil.JSONData(c, model.SupportEnvs)
}

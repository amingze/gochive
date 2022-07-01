package v1

import (
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/app/service"
	"github.com/amingze/gochive/internal/pkg/authed"
	"github.com/amingze/gochive/internal/pkg/bind"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
)

type AuthorizeResource struct {
	dAuthorize *dao.Authorize
	sAuthorize *service.Authorize
}

func NewAuthorizeResource() *AuthorizeResource {
	return &AuthorizeResource{
		dAuthorize: dao.NewAuthorize(),
		sAuthorize: service.NewAuthorize(),
	}
}

func (rs *AuthorizeResource) Register(router *gin.RouterGroup) {
	router.POST("/user/authorize", rs.create)              // 创建一个KEY
	router.GET("/user/authorize/:name", rs.find)           // 获取一个KEY
	router.PATCH("/user/authorize/:name/secret", rs.reset) // 重置KEY的secret
	router.DELETE("/user/authorize/:name", rs.remove)      // 重置KEY的secret
}

// create godoc
// @Tags UserAuthorize
// @Summary 创建秘钥
// @Description 创建秘钥
// @Accept json
// @Produce json
// @Param body body bind.BodyUserKeyCreation true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /user/authorize [post]
func (rs *AuthorizeResource) create(c *gin.Context) {
	p := new(bind.BodyUserKeyCreation)
	if err := c.ShouldBind(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	uk := model.NewAuthorize(authed.UidGet(c), p.Name)
	if err := rs.sAuthorize.Create(uk); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONData(c, uk)
}

// find godoc
// @Tags UserAuthorize
// @Summary 查询秘钥
// @Description 查询秘钥
// @Accept json
// @Produce json
// @Param name path string true "秘钥名称"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /user/authorize/{name} [get]
func (rs *AuthorizeResource) find(c *gin.Context) {
	uk, err := rs.dAuthorize.Find(authed.UidGet(c), c.Param("name"))
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	ginutil.JSONData(c, uk)
}

// reset godoc
// @Tags UserAuthorize
// @Summary 重置秘钥
// @Description 重置秘钥
// @Accept json
// @Produce json
// @Param name path string true "秘钥名称"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /user/authorize/{name}/secret [patch]
func (rs *AuthorizeResource) reset(c *gin.Context) {
	uk, err := rs.dAuthorize.Find(authed.UidGet(c), c.Param("name"))
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	if err := rs.sAuthorize.ResetSecret(uk); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSONData(c, uk)
}

// remove godoc
// @Tags UserAuthorize
// @Summary 删除秘钥
// @Description 删除秘钥
// @Accept json
// @Produce json
// @Param name path string true "秘钥名称"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /user/authorize/{name} [delete]
func (rs *AuthorizeResource) remove(c *gin.Context) {
	uk, err := rs.dAuthorize.Find(authed.UidGet(c), c.Param("name"))
	if err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}

	if err := rs.dAuthorize.Delete(uk); err != nil {
		ginutil.JSONServerError(c, err)
		return
	}

	ginutil.JSON(c)
}

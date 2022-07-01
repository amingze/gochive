package v1

import (
	"fmt"
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/app/service"
	"github.com/amingze/gochive/internal/pkg/authed"
	"github.com/amingze/gochive/internal/pkg/bind"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
	"strconv"
)

type MemoResource struct {
	sMemo *service.Memo
}

func NewMemoResource() *MemoResource {
	return &MemoResource{
		sMemo: service.NewMemo(),
	}
}

func (ms *MemoResource) Register(router *gin.RouterGroup) {
	router.GET("/memo/list", ms.findAll)
	router.GET("/memo/:id", ms.findOne)
	router.GET("/memo", ms.search)
	router.POST("/memo/backup", ms.backup)
	router.POST("/memo/backed", ms.backed)
	router.POST("/memo", ms.create)
	router.PUT("/memo/:id", ms.update)
	router.DELETE("/memo/:id", ms.delete)
}

// findAll godoc
// @Tags Memo
// @Summary 列出所有Memo
// @Description 列出所有Memo
// @Accept json
// @Produce json
// @Param query query bind.QueryMemo true "参数"
// @Success 200 {object} httputil.JSONResponse{data=gin.H{list=[]model.Memo,total=int64}}
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /memo/list [get]
func (ms *MemoResource) findAll(context *gin.Context) {
	p := new(bind.QueryMemo)
	if err := context.BindQuery(p); err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	uid := authed.UidGet(context)

	list, total, err := ms.sMemo.FindAll(uid, p.PageNo, p.PageSize)
	if err != nil {
		ginutil.JSONServerError(context, err)
		return
	}
	ginutil.JSONList(context, &list, total)
}

// findOne godoc
// @Tags Memo
// @Summary 查找Memo
// @Description 查找Memo
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} httputil.JSONResponse{data=gin.H{data=model.Memo}}
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /memo/{id} [get]
func (ms *MemoResource) findOne(context *gin.Context) {
	p := context.Param("id")
	id, err := strconv.ParseInt(p, 10, 64)
	if err != nil {
		ginutil.JSONServerError(context, err)
		return
	}
	memo, err := ms.sMemo.Find(id)
	if err != nil {
		ginutil.JSONServerError(context, err)
		return
	}
	ginutil.JSONData(context, memo)
}

// update godoc
// @Tags Memo
// @Summary 修改memo内容
// @Description 修改memo内容
// @Accept json
// @Produce json
// @Param id path string true "memo id"
// @Param body body bind.BodyMemoCreation true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /memo/{id} [put]
func (ms *MemoResource) update(context *gin.Context) {
	p := new(bind.BodyMemoCreation)
	if err := context.ShouldBindJSON(p); err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		ginutil.JSONBadRequest(context, err)
	}
	m := model.NewMemo()
	m.Content = p.Content
	m.Uid = authed.UidGet(context)
	if oldMemo, err := ms.sMemo.Find(id); oldMemo.Uid != authed.UidGet(context) || err != nil {
		ginutil.JSONBadRequest(context, err)
	}
	if err = ms.sMemo.Update(id, m); err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	ginutil.JSON(context)
}

// delete godoc
// @Tags Memo
// @Summary 删除一个Memo
// @Description 删除一个Memo
// @Accept json
// @Produce json
// @Param id path string true "Memo ID"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /memo/{id} [delete]
func (ms *MemoResource) delete(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	memo, err := ms.sMemo.Find(id)
	if err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	if memo == nil {
		ginutil.JSONBadRequest(context, fmt.Errorf("memo不存在"))
		return
	}
	if ms.sMemo.Delete(memo.ID); err != nil {
		ginutil.JSONServerError(context, err)
		return
	}
	ginutil.JSON(context)
}

// create godoc
// @Tags Memo
// @Summary 创建一个Memo
// @Description 创建一个Memo
// @Accept json
// @Produce json
// @Param body body bind.BodyMemoCreation true "参数"
// @Success 200 {object} httputil.JSONResponse{data=gin.H{data=model.Memo}}
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /memo [post]
func (ms *MemoResource) create(context *gin.Context) {
	p := new(bind.BodyMemoCreation)
	if err := context.BindJSON(p); err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	uid := authed.UidGet(context)
	m := model.NewMemo()
	m.Content = p.Content
	m.Uid = uid
	err := ms.sMemo.Create(m)
	if err != nil {
		ginutil.JSONServerError(context, err)
		return
	}
	ginutil.JSONData(context, m)
}

// search godoc
// @Tags Memo
// @Summary 查找Memo
// @Description 查找Memo
// @Accept json
// @Produce json
// @Param query query  bind.BodyMemoSearch true "参数"
// @Success 200 {object} httputil.JSONResponse{data=gin.H{list=[]model.Memo,total=int64}}
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /memo [get]
func (ms *MemoResource) search(context *gin.Context) {
	var p bind.BodyMemoSearch
	if err := context.ShouldBind(&p); err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	uid := authed.UidGet(context)

	list, total, err := ms.sMemo.Search(uid, p.Content, p.PageNo, p.PageSize)
	if err != nil {
		ginutil.JSONServerError(context, err)
		return
	}
	ginutil.JSONList(context, &list, total)
}

func (ms *MemoResource) backup(context *gin.Context) {
	uid := authed.UidGet(context)
	ms.sMemo.FindAll(uid, 0, 0)
}

func (ms *MemoResource) backed(context *gin.Context) {
	uid := authed.UidGet(context)
	ms.sMemo.FindAll(uid, 0, 0)

}

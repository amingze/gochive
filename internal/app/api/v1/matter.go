package v1

import (
	"github.com/amingze/gochive/internal/app/model"
	"github.com/amingze/gochive/internal/app/service"
	"github.com/amingze/gochive/internal/pkg/authed"
	"github.com/amingze/gochive/internal/pkg/bind"
	"github.com/amingze/gochive/internal/pkg/utils/encryptutil"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type FileResource struct {
	sFile   *service.File
	dMatter *model.Matter
	sMatter *service.Matter
}

func NewFileResource() ginutil.Resource {
	return &FileResource{
		sFile:   service.NewFile(),
		dMatter: &model.Matter{},
		sMatter: service.NewMatter(),
	}
}

func (rs *FileResource) Register(router *gin.RouterGroup) {
	router.POST("/matters", rs.fastload)
	router.POST("/matters/upload", rs.upload)
}

// fastload godoc
// @Tags Matters
// @Summary 快速上传
// @Description 根据文件信息上传,用于快速上传
// @Accept json
// @Produce json
// @Security OAuth2Application[matter, admin]
// @Param body body bind.BodyFastFile true "参数"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /matters [post]
func (rs *FileResource) fastload(c *gin.Context) {
	p := new(bind.BodyFastFile)
	if err := c.ShouldBindJSON(p); err != nil {
		ginutil.JSONBadRequest(c, err)
		return
	}
	uid := authed.UidGet(c)
	matter := model.NewMatter(uid, p.Name)
	if file, exist := rs.sFile.Exist(p.Signature, p.Size); exist {
		matter.Fid = file.ID
		matter.IsFast = true
		err := rs.sMatter.Create(matter)
		if err != nil {
			ginutil.JSONServerError(c, err)
			return
		}
	}

	ginutil.JSONData(c, matter)
}

// upload godoc
// @Tags Matters
// @Summary 上传文件
// @Description 上传文件
// @Accept json
// @Produce json
// @Security OAuth2Application[matter, admin]
// @Param file formData file true "文件"
// @Success 200 {object} httputil.JSONResponse
// @Failure 400 {object} httputil.JSONResponse
// @Failure 500 {object} httputil.JSONResponse
// @Router /matters/upload [post]
func (rs *FileResource) upload(context *gin.Context) {
	uid := authed.UidGet(context)
	file, err := context.FormFile("file")
	if err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	fileSignature := encryptutil.MD5(file)
	fileSize := file.Size
	fileModel, err := rs.sFile.Create(uid, fileSignature, fileSize)
	fileId := strconv.FormatInt(fileModel.ID, 10)
	matter := model.NewMatter(uid, file.Filename)
	matter.Fid = fileModel.ID
	matter.UpdatedAt = time.Now()
	context.SaveUploadedFile(file, fileId)
	rs.sMatter.Create(matter)
	if err != nil {
		ginutil.JSONBadRequest(context, err)
		return
	}
	ginutil.JSON(context)
}

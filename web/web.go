package web

import (
	"embed"
	"github.com/amingze/gochive/internal/pkg/middleware"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:embed app/dist/*
var embedFs embed.FS

type FileSystem struct {
	efs http.FileSystem
}

func NewFS() *FileSystem {
	return &FileSystem{http.FS(embedFs)}
}

func (fs FileSystem) Open(name string) (http.File, error) {
	f, err := fs.efs.Open(path.Join("app/dist", name))
	if os.IsNotExist(err) {
		return fs.efs.Open("app/dist/index.html")
	}
	return f, err
}
func SetupRoutes(ge *gin.Engine) {
	staticRouter := ge.Group("/")
	staticRouter.Use(gzip.Gzip(gzip.DefaultCompression))
	staticRouter.Use(middleware.Cors())
	ginutil.SetupEmbedAssets(staticRouter, NewFS(), "/js", "/css", "/font", "/img")
	ge.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/api") {
			return
		}
		c.FileFromFS(c.Request.URL.Path, NewFS())
	})
}

package ginutil

import (
	"github.com/amingze/gochive/internal/pkg/utils/httputil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"path"
	"strings"
)

type Resource interface {
	Register(router *gin.RouterGroup)
}

func SetupSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupResource(rg *gin.RouterGroup, resources ...Resource) {
	for _, resource := range resources {
		resource.Register(rg)
	}
}

func Startup(e *gin.Engine, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	go func() {
		logrus.Info("[rest server listen at ", srv.Addr, "]")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalln(err)
		}
	}()

	httputil.SetupGracefulStop(srv)
}

func SetupEmbedAssets(rg *gin.RouterGroup, fs http.FileSystem, relativePaths ...string) {
	handler := func(c *gin.Context) {
		c.FileFromFS(strings.TrimPrefix(c.Request.URL.Path, rg.BasePath()), fs)
	}

	for _, relativePath := range relativePaths {
		urlPattern := relativePath
		if urlPattern != "/" {
			urlPattern = path.Join(relativePath, "/*filepath")
		}
		rg.GET(urlPattern, handler)
		rg.HEAD(urlPattern, handler)
	}
}

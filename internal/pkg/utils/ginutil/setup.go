package ginutil

import (
	"github.com/amingze/gochive/internal/pkg/utils/httputil"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
)

type Resource interface {
	Register(router *gin.RouterGroup)
}

func SetupResource(rg *gin.RouterGroup, resources ...Resource) {
	for _, resource := range resources {
		resource.Register(rg)
	}
}

func SetupPing(e *gin.Engine) {
	pingHandler := func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}

	e.HEAD("/ping", pingHandler)
	e.GET("/ping", pingHandler)
}

func SetupSwagger(engine *gin.Engine) {
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Startup(e *gin.Engine, addr string) {
	srv := &http.Server{
		Addr:    addr,
		Handler: e,
	}

	go func() {
		log.Printf("[rest server listen at %s]", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	httputil.SetupGracefulStop(srv)
}

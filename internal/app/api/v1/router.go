package v1

import (
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"github.com/gin-gonic/gin"

	_ "github.com/amingze/gochive/internal/docs"
)

// @title gochive
// @description gochive apis
// @version 0.0.1

// @BasePath /api/
// @securitydefinitions.oauth2.application OAuth2Application
// @scope.matter Grants matter access and write
// @scope.admin Grants read and write access to administrative information
// @tokenUrl /api/tokens
// @name Authorization

// @contact.name API Support
// @contact.url https://github.com/amingze/gochive
// @contact.email

// @license.name Apache 2.0
// @license.url https://github.com/amingze/gochive/blob/master/LICENSE

func SetupRoutes(ge *gin.Engine) {
	ginutil.SetupSwagger(ge)

	apiRouter := ge.Group("/api")
	ginutil.SetupResource(apiRouter,
		NewOptionResource(),
		NewUserResource(),
		NewAuthorizeResource(),
		NewTokenResource(),
		NewFileResource(),
		NewMemoResource(),
	)
}

package middleware

import (
	"fmt"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Installer(c *gin.Context) {
	if !viper.IsSet("installed") {
		ginutil.JSONError(c, 520, fmt.Errorf("系统未初始化"))
		return
	}
}

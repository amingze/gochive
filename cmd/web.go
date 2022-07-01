package cmd

import (
	"fmt"
	"github.com/amingze/gochive/internal/app/api/v1"
	"github.com/amingze/gochive/internal/app/dao"
	"github.com/amingze/gochive/internal/pkg/utils/ginutil"
	"github.com/amingze/gochive/web"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "web",
	Short: "A cloud disk base on the cloud service.",
	Run: func(cmd *cobra.Command, args []string) {
		serverRun()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntP("port", "p", 8222, "server port")
	logrus.Debug("传入端口", viper.GetInt("port"))

	viper.BindPFlags(serverCmd.Flags())
}

func serverRun() {
	gin.SetMode(gin.DebugMode)
	if viper.IsSet("installed") {
		logrus.Debug("应用已安装，加载数据库", viper.GetString("database.driver"), viper.GetString("database.dsn"))
		dao.Init(viper.GetString("database.driver"), viper.GetString("database.dsn"))
	}

	ge := gin.Default()
	v1.SetupRoutes(ge)
	web.SetupRoutes(ge)
	logrus.Debug("端口", viper.GetInt("port"))

	addr := fmt.Sprintf(":%d", viper.GetInt("port"))
	ginutil.Startup(ge, addr)
}

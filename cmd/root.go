package cmd

import (
	"fmt"
	"github.com/amingze/gochive/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd
var rootCmd = &cobra.Command{
	Use:   "gochive",
	Short: "gochive is a archive application.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Debug("异常退出")
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is "+config.DefaultConfigFile+")")
}

// initConfig
func initConfig() {
	logrus.Debug("加载配置文件")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(config.DefaultConfigFilePath)
		viper.SetConfigName(config.DefaultConfigFileName)
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Debug("使用配置文件:", viper.ConfigFileUsed())
		viper.Set("installed", true)
	} else {
		logrus.Error(err)
	}
}

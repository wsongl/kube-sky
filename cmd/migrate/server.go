package migrate

import (
	"fmt"
	"kube-sky/pkg/conn"
	"kube-sky/pkg/logger"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configYml string
	structure bool
	StartCmd  = &cobra.Command{
		Use:          "migrate",
		Short:        "Synchronous data structure",
		Example:      "kube-sky migrate -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/settings.yml", "指定配置文件启动服务")
	StartCmd.PersistentFlags().BoolVarP(&structure, "structure", "s", false, "仅同步数据结构")
}

func setup() {
	// 加载配置文件
	viper.SetConfigFile(configYml) // 指定配置文件
	err := viper.ReadInConfig()    // 读取配置信息
	if err != nil {                // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 日志配置
	logger.Setup()

	// 数据库配置
	conn.Setup()
}

func run() (err error) {
	// 同步数据结构
	AutoMigrate(structure)

	// 同步初始数据
	return
}

package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config") // 指定配置文件名称，不需要带后缀
	viper.SetConfigType("yaml")   // 指定配置文件的类型
	viper.AddConfigPath(".")      // 指定查找配置文件的路径，这里指使用相对路径
	if err = viper.ReadInConfig(); err != nil {
		fmt.Printf("viper.ReadInConfig() failed: %s", err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件修改了......")
	})
	return
}

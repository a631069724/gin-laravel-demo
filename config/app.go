package Config

import (
	"github.com/spf13/viper"
	"log"
)

var App *viper.Viper

//初始化config数据
func InitConfig(env string) {
	//viper 读取配置文件参数
	App = viper.New()

	App.AddConfigPath("./config")
	//读取dev配置文件.如果设计配置文件的切换请在提交前修改
	//TODO:这里是手动切换配置文件,没有想到更好的方法
	App.SetConfigType("toml")
	App.SetConfigName("app" + "." + env)
	err := App.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

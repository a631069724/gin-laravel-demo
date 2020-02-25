package Config

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

var (
	Env       string       // 环境
	AppPath   string       // 程序运行目录
	App       *viper.Viper // 配置
	Db        *gorm.DB     // 数据库engine
	RedisPool *redis.Pool
)

//初始化config数据
func InitConfig(env string) {
	ep, err := ExecPath()
	if err != nil {
		panic("获取程序运行文件失败: " + err.Error())
	}
	log.Printf("程序运行文件路径: %s", ep)
	ep = strings.Replace(ep, "\\", "/", -1)
	AppPath = filepath.Dir(ep)
	log.Printf("程序运行文件目录: %s", AppPath)

	//viper 读取配置文件参数
	App = viper.New()

	App.AddConfigPath("./config")
	//读取dev配置文件.如果设计配置文件的切换请在提交前修改
	//TODO:这里是手动切换配置文件,没有想到更好的方法
	App.SetConfigType("toml")
	App.SetConfigName("app" + "." + env)
	err = App.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

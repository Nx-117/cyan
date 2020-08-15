package config

/**
读取config相关
*/
import (
	"github.com/spf13/viper"
)

var LoadC *viper.Viper

/**
以下三个参数需要自己设置
才可以使用次工具类
配置文件名字
*/
var Name = "config"

/**
配置文件路径
*/
var Path = "config"

/**
配置文件类型
*/
var Type = "yaml"

func LoadConfig() {
	v := viper.New()
	v.SetConfigName(Name) //文件名
	v.AddConfigPath(Path) // 路径
	//windows环境下为%GOPATH，linux环境下为$GOPATH
	//v.AddConfigPath("%GOPATH/src/")
	//设置配置文件类型
	v.SetConfigType(Type)

	if err := v.ReadInConfig(); err != nil {
		panic("load config error: " + err.Error())
	}
	LoadC = v
}

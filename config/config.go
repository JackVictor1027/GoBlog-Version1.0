package config

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig
	//名称首字母大写，表示可以提供给外部进行访问
}

// 模板数据并不导入到数据库当中，而是与配置文件连接
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}

type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecret     string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig

// 习惯使用初始化方法
func init() {
	Cfg = new(tomlConfig)

	Cfg.System.AppName = "GoBlog"
	Cfg.System.Version = 1.0
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	Cfg.System.CurrentDir = currentDir
	_, err = toml.DecodeFile("config/config.toml", &Cfg)
	if err != nil {
		panic(err) //程序直接停止
	}
}

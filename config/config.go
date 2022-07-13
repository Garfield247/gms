package config

import (
	"log"

	"github.com/go-ini/ini"
)

type Path struct {
	ImgSavePath string
	Img404Path  string
}

type Server struct {
	Port    int
	RunMode string
}

var PathSetting = &Path{}
var ServerSetting = &Server{}

func init() {
	log.Println("初始化配置")
	Cfg, err := ini.Load("./gmsconfig.ini")
	if err != nil {
		log.Fatalf("未找到文件`./gmsconfig.ini` :%v", err)
	} else {
		log.Println("配置文件读取成功")
	}

	err = Cfg.Section("path").MapTo(PathSetting)
	if err != nil {
		log.Fatalf("解析路径配置失败:%v", err)
	}
	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("解析服务配置失败:%v", err)
	}
	log.Println("配置初始化完成")
}

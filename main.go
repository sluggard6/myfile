package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/sluggard/myfile/application"
	"github.com/sluggard/myfile/config"
	_ "github.com/sluggard/myfile/docs"
)

var configFile = flag.String("c", config.DefaultConfigPath, "配置文件路径")
var version = flag.Bool("v", false, "打印版本号")

//Version 程序版本号
const Version = "0.0.1"

func main() {
	log.SetLevel(log.TraceLevel)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [options] [command]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "  -c <path>\n")
		fmt.Fprintf(os.Stderr, "    设置项目配置文件路径，可选\n")
		fmt.Fprintf(os.Stderr, "  -v 打印项目版本号，默认为: false\n")
		fmt.Fprintf(os.Stderr, "    打印版本号\n")
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.Parse()

	if *version {
		fmt.Println(fmt.Sprintf("版本号：%s\n", Version))
		return
	}

	irisServer := application.NewServer(config.LoadConfing(*configFile))
	log.Info(config.GetConfig())
	if irisServer == nil {
		panic("http server 初始化失败")
	}

	// if libs.IsPortInUse(libs.Config.Port) {
	// 	if !irisServer.Status {
	// 		panic(fmt.Sprintf("端口 %d 已被使用\n", libs.Config.Port))
	// 	}
	// 	irisServer.Stop() // 停止
	// }

	err := irisServer.Start()
	if err != nil {
		panic(fmt.Sprintf("http server 启动失败: %+v", err))
	}

	// logging.InfoLogger.Infof("http server %s:%d start", libs.Config.Host, libs.Config.Port)

}

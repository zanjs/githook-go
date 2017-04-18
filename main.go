package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/configor"
)

var port int

// 初始化参数
func init() {

	configor.Load(&Config, "config.yml")

	var welcomeStr = "\n欢迎使用 git 微服务 \n- 高性能、跨平台、易扩展 \n- Version: " + Config.APPInfo.Version
	var portRed = "\n- 服务占用 " + strconv.Itoa(Config.APPInfo.Port) + " 端口号"
	var ymlStr = "\n\nconfig.yml 文件为自定义配置文件, 您可以根据场景修改它"
	var openStr = welcomeStr + portRed + ymlStr + "\n\n "

	flag.IntVar(&port, "port", Config.APPInfo.Port, "服务器端口")
	// fmt.Println(openStr + fmt.Sprintf("%d", port))
	fmt.Println(openStr)
	flag.Parse()

}

func main() {

	http.HandleFunc("/pull", Pull)

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

var (
	logFileName = flag.String("log", "log.log", "Log file name")
)

package main

import (
	"shortlib"
	"fmt"
	"flag"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "conf", "config.ini", "configure file full path")
	flag.Parse()

	//读取配置文件
	fmt.Printf("[INFO] Read configure file...\n")
	configure, err := shortlib.NewConfigure(configFile)
	if err != nil {
		fmt.Printf("[ERROR] Parse Configure File Error: %v\n", err)
		return
	}

	redis_cli, err := shortlib.NewRedisAdaptor(configure)
	if err != nil {
		fmt.Printf("[ERROR] Redis init fail..\n")
		return
	}
	redis_cli.SetUrl("https://kdt.im/W7CBgr", "https://wap.koudaitong.com/v2/showcase/homepage?alias=edm86aq2")

	original_url, err := redis_cli.GetUrl("https://kdt.im/W7CBgr")
	if err != nil {
		fmt.Printf("[ERROR] Parse Configure File Error: %v\n", err)
	}

	fmt.Printf(original_url)
}
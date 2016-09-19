package main

import (
	"shortlib"
	"fmt"
	"flag"
	"net/http"
	"os"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "conf", "./config.ini", "configure file full path")
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
	if configure.GetRedisStatus() {
		err = redis_cli.InitCountService()
		if err != nil {
			fmt.Printf("[ERROR] Init Redis key count fail...\n")
		}
	}

	//不使用redis的情况
	count_channl := make(chan shortlib.CountChannl, 1000)
	go CountThread(count_channl)
	count_type := configure.GetCounterType()

	countfunction := shortlib.CreateCounter(count_type, count_channl, redis_cli)

	//启动LRU缓存
	fmt.Printf("[INFO] Start LRU Cache System...\n")
	lru, err := shortlib.NewLRU(redis_cli)
	if err != nil {
		fmt.Printf("[ERROR]LRU init fail...\n")
	}

	//初始化两个短连接服务
	fmt.Printf("[INFO] Start Service...\n")
	baseprocessor := &shortlib.BaseProcessor{redis_cli, configure, configure.GetHostInfo(), lru, countfunction}

	original := &OriginalProcessor{baseprocessor, count_channl}
	short := &ShortProcessor{baseprocessor}

	//启动http handler
	router := &shortlib.Router{configure, map[int]shortlib.Processor{
		0: short,
		1: original,
	}}

	//启动服务

	port, _ := configure.GetPort()
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("[INFO]Service Starting addr :%v,port :%v\n", addr, port)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		os.Exit(1)
	}
}
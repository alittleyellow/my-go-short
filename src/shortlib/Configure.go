package shortlib

import (
	"fmt"
	"errors"
	"github.com/ewangplay/config"
	"strconv"
)

type Configure struct {
	ConfigureMap map[string]string
}

func NewConfigure(fileName string) (*Configure, error) {
	config := &Configure{}
	config.ConfigureMap = make(map[string]string)
	err := config.ParseConfigure(fileName)
	if (err) != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return config, nil
}

func (this *Configure) ParseConfigure(fileName string) error {
	cfg, err := config.ReadDefault(fileName);

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	this.loopConfigure("server", cfg)
	this.loopConfigure("service", cfg)
	this.loopConfigure("redis", cfg)
	return nil
}

func (this *Configure) loopConfigure(sectionName string, cfg *config.Config) error {
	if cfg.HasSection(sectionName) {
		section, err := cfg.SectionOptions(sectionName)
		if err == nil {
			for _, v := range section {
				options, err := cfg.String(sectionName, v)
				if err == nil {
					this.ConfigureMap[v] = options
				}
			}
			return nil 
		}
		return errors.New("Parse Error")
	}
	return errors.New("Parse Error")
}

//服务信息
func (this *Configure) GetPort() (int, error) {

	portstr, ok := this.ConfigureMap["port"]
	if ok == false {
		return 9090, errors.New("No Port set, use default")
	}

	port, err := strconv.Atoi(portstr)
	if err != nil {
		return 9090, err
	}

	return port, nil
}

func (this *Configure) GetRedisHost() (string, error) {
	redishost, ok := this.ConfigureMap["redishost"]

	if ok == false {
		return "127.0.0.1", errors.New("No redishost,use defualt")
	}

	return redishost, nil
}

func (this *Configure) GetRedisPort() (string, error) {
	redisport, ok := this.ConfigureMap["redisport"]

	if ok == false {
		return "6379", errors.New("No redisport,use defualt")
	}

	return redisport, nil
}

func (this *Configure) GetRedisStatus() bool {

	status, ok := this.ConfigureMap["status"]
	if ok == false {
		return true
	}

	if status == "true" {
		return true
	}
	return false

}

func (this *Configure) GetHostInfo() string {

	host_name, ok := this.ConfigureMap["hostname"]
	if ok == false {
		return "http://wusay.org/"
	}

	return host_name

}

func (this *Configure) GetCounterType() string {

	count_type, ok := this.ConfigureMap["counter"]
	if ok == false {
		return "inner"
	}

	return count_type

}
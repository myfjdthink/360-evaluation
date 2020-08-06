package config

import (
	"fmt"
	"github.com/jinzhu/configor"
)

var Config = struct {
	Port int `default:"3000"`
}{}

func init() {

	err := configor.Load(&Config, "config/config.yaml")
	if err != nil {
		fmt.Errorf("load Config Error %e", err)
		return
	}
	fmt.Println("config Port", Config.Port)
}

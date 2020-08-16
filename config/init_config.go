package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"time"
)

type JWT struct {
	Secret string
	Expire time.Duration
}

var Config = struct {
	Port     int `default:"3000"`
	Password string
	JWT      JWT
}{}

func init() {

	err := configor.Load(&Config, "config/config.yaml")
	if err != nil {
		err = fmt.Errorf("load Config Error %e", err)
		return
	}
	fmt.Println("config Port", Config.Port)
}

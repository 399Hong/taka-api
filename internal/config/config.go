package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Mysql struct { //more config
		DataSource string
	}
	Auth struct {
		AccessSecret string
		ExpiresInSec int64
	}
}

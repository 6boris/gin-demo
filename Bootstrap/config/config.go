package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	SecretKey  string
	EncryptKey string
}

var AppConfig = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	HttpAddr     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}

type MySQL struct {
	Debug       bool
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var MySQLConfig = &MySQL{}

type Jwt struct {
	Secret string
}

var JwtConfig = &Jwt{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("env.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'env.ini': %v", err)
	}
	mapTo("app", AppConfig)
	mapTo("server", ServerConfig)
	mapTo("mysql", MySQLConfig)
	mapTo("jwt", JwtConfig)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}

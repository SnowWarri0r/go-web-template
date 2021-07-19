package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Load struct {
	Server Server `yaml:"server"`
	Dev    Conf   `yaml:"dev"`
	Prod   Conf   `yaml:"prod"`
}
type Config struct {
	Server Server
	Conf   Conf
}
type Server struct {
	Addr string `yaml:"addr"`
	Salt string `yaml:"salt"`
}
type Conf struct {
	DB                      string `yaml:"db"`
	MySQLMaxOpenConnections int    `yaml:"mysqlMaxOpenConns"`
	MySQLMaxIdleConnections int    `yaml:"mysqlMaxIdleConns"`
	MySQLConnMaxIdleTime    int64  `yaml:"mysqlConnMaxIdleTime"`
	MySQLConnMaxLifeTime    int64  `yaml:"mysqlConnMaxLifeTime"`
}

var load = &Load{}
var config = &Config{}

func InitConfig() {
	f, err := os.Open("./config/config.yaml")
	if err != nil {
		Log().Println(err.Error())
		Log().Fatalln("Opening config file failed")
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		Log().Fatalln("Reading config file failed")
	}
	err = yaml.UnmarshalStrict(data, &load)
	if err != nil {
		Log().Fatalln("Config has format error")
	}
}

var (
	mode          string
	mysqlUser     string
	mysqlPasswd   string
	mysqlDatabase string
)

func LoadConfig() {
	config.Server = load.Server
	mode = os.Getenv("APP_MODE")
	mysqlUser = os.Getenv("MYSQL_USER")
	mysqlPasswd = os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	if len(mode) < 1 {
		mode = "dev"
		config.Conf = load.Dev
	} else if mode == "prod" {
		config.Conf = load.Prod
	} else {
		config.Conf = load.Dev
	}
}

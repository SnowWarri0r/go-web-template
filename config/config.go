package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var config = make(map[interface{}]interface{})

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
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		Log().Fatalln("Config has format error")
	}
}

var (
	addr          string
	salt          string
	mode          string
	mysqlUser     string
	mysqlPasswd   string
	mysqlDatabase string
	dbUrl         string
	mysqlMaxOpenConns int
	mysqlMaxIdleConns int
	mysqlConnMaxIdleTime int
	mysqlConnMaxLifeTime int
)

func LoadConfig() {
	var ok bool
	mode = os.Getenv("APP_MODE")
	mysqlUser = os.Getenv("MYSQL_USER")
	mysqlPasswd = os.Getenv("MYSQL_PASSWORD")
	mysqlDatabase = os.Getenv("MYSQL_DATABASE")
	if len(mode) < 1 {
		mode = "dev"
	}
	if addr, ok = config["server"].(map[interface{}]interface{})["addr"].(string); !ok {
		Log().Fatalln("cannot read addr config")
	}
	if salt, ok = config["server"].(map[interface{}]interface{})["salt"].(string); !ok {
		Log().Fatalln("cannot read salt config")
	}
	if mode == "prod" {
		if dbUrl, ok = config["prod"].(map[interface{}]interface{})["db"].(string); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlMaxOpenConns, ok = config["prod"].(map[interface{}]interface{})["mysqlMaxOpenConns"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlMaxIdleConns, ok = config["prod"].(map[interface{}]interface{})["mysqlMaxIdleConns"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlConnMaxIdleTime, ok = config["prod"].(map[interface{}]interface{})["mysqlConnMaxIdleTime"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlConnMaxLifeTime, ok = config["prod"].(map[interface{}]interface{})["mysqlConnMaxLifeTime"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
	} else {
		if dbUrl, ok = config["dev"].(map[interface{}]interface{})["db"].(string); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlMaxOpenConns, ok = config["dev"].(map[interface{}]interface{})["mysqlMaxOpenConns"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlMaxIdleConns, ok = config["dev"].(map[interface{}]interface{})["mysqlMaxIdleConns"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlConnMaxIdleTime, ok = config["dev"].(map[interface{}]interface{})["mysqlConnMaxIdleTime"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
		if mysqlConnMaxLifeTime, ok = config["dev"].(map[interface{}]interface{})["mysqlConnMaxLifeTime"].(int); !ok {
			Log().Fatalln("cannot read db config")
		}
	}
}

package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

var config=make(map[interface{}]interface{})

func InitConfig() {
	f,err:=os.Open("./config/config.yaml")
	if err != nil {
		log.Println(err.Error())
		log.Fatalln("Opening config file failed")
	}
	data,err:=ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln("Reading config file failed")
	}
	err=yaml.Unmarshal(data,&config)
	if err != nil {
		log.Fatalln("Config has format error")
	}
}
var (
	addr string
	salt string
)
func LoadConfig() {
	var ok bool
	if addr,ok=config["server"].(map[interface{}]interface{})["addr"].(string);!ok{
		log.Fatalln("cannot read addr config")
	}
	if salt,ok=config["server"].(map[interface{}]interface{})["salt"].(string);!ok{
		log.Fatalln("cannot read salt config")
	}
}
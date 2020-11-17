package core

import (
	"io/ioutil"
	"os"
)
import "gopkg.in/yaml.v2"

func loadConfig(root_path string) Monitor_yaml {
	//load config
	//str,_ := os.Getwd()
	var b []byte
	var monitoryaml Monitor_yaml
	yaml_file := root_path + "/config.yaml"
	Debugf(" yaml path:%s\n", yaml_file)
	b, err := ioutil.ReadFile(yaml_file)
	if err != nil {
		Errorf("Can't load file: config.yaml%s\n", err)
		os.Exit(-1)
	}

	err2 := yaml.Unmarshal(b, &monitoryaml)
	if err != nil {
		Errorf("%s\n", err2)
	}
	var serverApi = monitoryaml.Monitor.ServerJiang.ServerJiangApi
	Debugf(" check api:%s\n", serverApi)

	if !CheckYaml(monitoryaml) {
		Errorf("something wrong with the config.yaml")
		os.Exit(-1)
	}
	return monitoryaml
}
func CheckYaml(monitorYaml Monitor_yaml) bool {
	//todo check action
	return true
}

// db/config.go
// 設定ファイルを読み込み、構造体にマッピングを行う

package db

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

func LoadConfig(path string) Config {
	var config Config
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("config.yml read error: %v", err)
	}
	if err := yaml.Unmarshal(configFile, &config); err != nil {
		log.Fatalf("YAML unmarshal error: %v", err)
	}
	return config
}

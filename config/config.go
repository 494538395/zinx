package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const configPath = "config/config.yaml"

func InitConfig() error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &Config)
}

var Config ConfigStruct

type ConfigStruct struct {
	Server struct {
		ServerName       string `yaml:"serverName"`
		IP               string `yaml:"IP"`
		Port             int    `yaml:"port"`
		IPVersion        string `yaml:"IPVersion"`
		MaxWorkerTaskLen uint32 `yaml:"maxWorkerTaskLen"`
		WorkerPoolSize   uint32 `yaml:"workerPoolSize"`
		MaxConnSize      uint32 `yaml:"maxConnSize"` // 服务器最大连接数
	} `yaml:"server"`
}

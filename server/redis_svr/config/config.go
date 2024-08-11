package config

import (
	"os"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

var (
	fileName = "grpc.yaml"
)

type Config struct {
	ServerConfig ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Service ServiceConfig `yaml:"service"`
}

type ServiceConfig struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Port    int
	IP      string
}

var globalConfig Config

func init() {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &globalConfig)
	if err != nil {
		panic(err)
	}
	strs := strings.Split(globalConfig.ServerConfig.Service.Address, ":")
	if len(strs) != 2 {
		panic("address error")
	}
	globalConfig.ServerConfig.Service.IP = strs[0]
	port, err := strconv.Atoi(strs[1])
	if err != nil {
		panic(err)
	}
	globalConfig.ServerConfig.Service.Port = port
}

func GetPort() int {
	return globalConfig.ServerConfig.Service.Port
}

func GetAddress() string {
	return globalConfig.ServerConfig.Service.Address
}

func GetServiceName() string {
	return globalConfig.ServerConfig.Service.Name
}

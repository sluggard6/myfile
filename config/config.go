package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Server   Server
	Database Database
	Stroe    Store
}

type Server struct {
	Host string
	Port int
}

type Store struct {
	DataRoot string
}

type Database struct {
	Type     dbType
	Url      string
	Username string
	Password string
}

type dbType string
type fileType string

const (
	DefaultConfigPath string   = "application.json"
	Mysql             dbType   = "mysql"
	Sqlite            dbType   = "sqlite"
	Yml               fileType = "yml"
	Json              fileType = "json"
)

var config Config

func GetConfig() *Config {
	return &config
}

func New(config string) Config {
	// return Config{"127.0.0.1", 5678}
	return LoadConfing(DefaultConfigPath)
}

func LoadConfing(path string) Config {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("unable to decode into struct, %v", err)
		// panci(err)
		// return nil, err
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}

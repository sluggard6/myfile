package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Server Server
	Db     Db
}

type Server struct {
	Host string
	Port int
}

type Db struct {
	DbType   string
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

func New(config string) Config {
	// return Config{"127.0.0.1", 5678}
	return LoadConfing(DefaultConfigPath)
}

func LoadConfing(path string) Config {
	var config Config
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

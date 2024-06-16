package conf

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"
)

type UserServConf struct {
	Port    int  `json:"port"`
	OpenTls bool `json:"open_tls"`
}

type MysqlConf struct {
	Addr    string `json:"addr"`
	User    string `json:"user"`
	Pwd     string `json:"pwd"`
	DbName  string `json:"db_name"`
	LogMode bool   `json:"log_mode"`
}

type Conf struct {
	UserServConf UserServConf `json:"user_serv"`
	MysqlConf    MysqlConf    `json:"mysql"`
}

var cfg Conf

func InitConfig() error {
	jsonFile, err := os.Open(configPath)
	if err != nil {
		return errors.New("can't find config in path : " + configPath)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &cfg)
	return err
}

func GetSqlConf() MysqlConf {
	return cfg.MysqlConf
}

func GetUserServConf() UserServConf {
	return cfg.UserServConf
}

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config_path", "./conf/config.json", "please input config path")
}

func InitEnv() {
	flag.Parse()
}

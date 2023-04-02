package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Conf struct {
	Mysql MySqlConfig `yaml:"mysql" json:"mysql"`
	Redis RedisConfig `yaml:"redis" json:"redis"`
	Minio MinioConfig `yaml:"minio" json:"minio"`
}

type MySqlConfig struct {
	Url    string `yaml:"url" json:"url"`
	Port   uint16 `yaml:"port" json:"port"`
	User   string `yaml:"user" json:"user"`
	Passwd string `yaml:"passwd" json:"passwd"`
	DbName string `yaml:"dbName" json:"dbName"`
}

type RedisConfig struct {
	Addr   string `yaml:"addr" json:"addr"`
	Passwd string `yaml:"passwd" json:"passwd"`
}

type MinioConfig struct {
	IP              string `yaml:"ip" json:"ip"`
	Port            uint16 `yaml:"port" json:"port"`
	AccessKeyID     string `yaml:"accessKeyID" json:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey" json:"secretAccessKey"`
}

func ReloadConfig() (*Conf, error) {
	return loadConfig()
}

func LoadConfig(fileName string, ch chan struct{}) (*Conf, error) {
	return initConfig(fileName, ch)
}

func initConfig(fileName string, ch chan struct{}) (*Conf, error) {
	viper.SetConfigFile(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		ch <- struct{}{}
	})
	return loadConfig()
}

func loadConfig() (*Conf, error) {
	var conf *Conf
	err := viper.Unmarshal(&conf)
	return conf, err
}

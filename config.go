package dream

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ParseConfig(configFile string) *Config {
	fileBytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	serverCfg := new(Config)
	err = yaml.Unmarshal(fileBytes, serverCfg)
	if err != nil {
		panic(err)
	}
	return serverCfg
}

type Config struct {
	Server *Server `yaml:"server"`
	Mysql  *Mysql  `yaml:"mysql"`
	Redis  *Redis  `yaml:"redis"`
	Aliyun *Aliyun `yaml:"aliyun"`
	Mail   *Mail   `yaml:"mail"`
}

type Server struct {
	Address string `yaml:"address"`
	Port    string `yaml:"port"`
}

type Mysql struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
}

type Redis struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type Aliyun struct {
	Sms *Sms `yaml:"sms"`
}

type Sms struct {
	Url             string `yaml:"url"`
	AccessKeyId     string `yaml:"access-key-id"`
	AccessKeySecret string `yaml:"access-key-secret"`
	SignName        string `yaml:"sign-name"`
	TemplateCode    string `yaml:"template-code"`
}

type Mail struct {
	SenderMail string `yaml:"sender-mail"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Smtp       string `yaml:"smtp"`
	Port       string `yaml:"port"`
}

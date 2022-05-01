package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Server holds data necessary for server configuration
type Server struct {
	Port         string `yaml:"port"`
	Debug        bool   `yaml:"debug"`
	ReadTimeout  int    `yaml:"read_timeout_seconds"`
	WriteTimeout int    `yaml:"write_timeout_seconds"`
}

type RabbitMq struct {
	QueueName       string `yaml:"queue_name"`
	RabbitMQConnUrl string `yaml:"rabbit_mq_conn_url"`
}

type Configuration struct {
	Server   Server   `yaml:"server"`
	DbPSN    string   `yaml:"db_psn"`
	RabbitMQ RabbitMq `yaml:"rabbitmq"`
}

func LoadConfigs(path string) (*Configuration, error) {
	configs, err := loadFromFile(path)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func loadFromFile(path string) (*Configuration, error) {
	fmt.Printf("loading config file: %s\n", path)
	var cfg = new(Configuration)
	var files []string

	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	if err := yaml.Unmarshal(bytes, cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return cfg, err
}

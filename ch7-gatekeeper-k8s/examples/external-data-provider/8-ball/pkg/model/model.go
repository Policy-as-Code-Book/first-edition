package model

import (
	"errors"

	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"jimmyray.io/magic-8-ball/pkg/utils"
	"math/rand"
	"time"
)

const (
	ContentTypeGzip   string = "application/gzip"
	ContentTypeJson   string = "application/json"
	HeaderContentType string = "Content-Type"
	PageNotFound      string = "404 page not found"
	Forbidden         string = "403 forbidden"
	Unauthorized      string = "401 unauthorized"

	timeout    = 1 * time.Second
	apiVersion = "externaldata.gatekeeper.sh/v1alpha1"
	kind       = "ProviderResponse"
)

type Answer struct {
	Answer string `yaml:"answer" json:"answer"`
	Code   string `yaml:"code" json:"code"`
}

type Config struct {
	Answers []Answer `yaml:"answers"`
	Network struct {
		ServerAddress string `yaml:"serverAddress"`
		ServerPort    string `yaml:"serverPort"`
		TLS           struct {
			Enabled bool   `yaml:"enabled"`
			Cert    string `yaml:"cert"`
			Key     string `yaml:"key"`
		} `yaml:"tls"`
	} `yaml:"network"`
}

var (
	AC         Config
	ConfigFile string
)

func (c *Config) Yaml() ([]byte, error) {
	out, err := yaml.Marshal(&c)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Config) Json() ([]byte, error) {
	out, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// LoadConfig Load config from provided file
func (c *Config) LoadConfig(configFile string) error {
	if !utils.FileExists(configFile) {
		err := errors.New(configFile + " does not exist")
		return err
	}

	bytes, err := utils.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) getAnswer() *Answer {
	r := c.Answers[rand.Intn(len(c.Answers))]
	return &r
}

func getAnswerJson() (string, error) {
	m := AC.getAnswer()
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(j), nil
}

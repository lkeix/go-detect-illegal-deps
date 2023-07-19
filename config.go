package detectillegaldeps

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	basePath       string
	internalPrefix string
	whiteList      map[string][]string
}

func NewConfig(yamlPath, basePath string) (*Config, error) {
	y, err := parseYaml(yamlPath)
	if err != nil {
		return nil, err
	}
	return &Config{
		basePath:       basePath,
		internalPrefix: y.InternalPrefix,
		whiteList:      y.WhiteList,
	}, nil
}

type Yaml struct {
	InternalPrefix string              `yaml:"internalPrefix"`
	WhiteList      map[string][]string `yaml:"whitelist"`
}

func parseYaml(path string) (*Yaml, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	y := new(Yaml)
	err = yaml.Unmarshal(b, y)
	return y, err
}

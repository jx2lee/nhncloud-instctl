package instance

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// YamlParser parser
type YamlParser struct {
	Version string `yaml:"version"`
	Spec    struct {
		Credentials struct {
			TenantID string `yaml:"tenantID"`
			Username string `yaml:"accessID"`
			Password string `yaml:"accessPassword"`
		} `yaml:"credentials"`
		PrivateKeys []struct {
			Name string `yaml:"keyName"`
			Path string `yaml:"keyPath"`
		} `yaml:"privateKeys"`
	}
}

// Envparser have config info
type Envparser struct {
	configPath  string
	yamlContent YamlParser
}

// GetEnvparser get parser object
func GetEnvparser(args ...string) *Envparser {
	var configPath string
	if len(args) > 0 {
		configPath = "/etc/toast-rds-config.yaml"
	} else {
		configPath = "/etc/toast-config.yaml"
	}
	var envParser = Envparser{configPath: configPath}
	envParser.OpenConfigFile()
	return &envParser
}

// OpenConfigFile open NHN Cloud config in /etc/toast-config.yaml
func (ep *Envparser) OpenConfigFile() (result map[string]interface{}) {
	yamlFile, err := ioutil.ReadFile(ep.configPath)
	if err != nil {
		logrus.Fatal("Unable to open /etc/toast-config.yaml")
	}
	err = yaml.Unmarshal(yamlFile, &ep.yamlContent)
	return result
}

// GetPasswordCredentials get Username/Passwd
func (ep *Envparser) GetPasswordCredentials() (string, string, string, error) {
	tennantID := ep.yamlContent.Spec.Credentials.TenantID
	accessID := ep.yamlContent.Spec.Credentials.Username
	accessPassword := ep.yamlContent.Spec.Credentials.Password
	return tennantID, accessID, accessPassword, nil
}

// GetPrivateKeyPath return private key path
func (ep *Envparser) GetPrivateKeyPath(keyName string) string {
	var keyPath string
	for _, value := range ep.yamlContent.Spec.PrivateKeys {
		if value.Name == keyName {
			keyPath = value.Path
		}
	}
	return keyPath
}

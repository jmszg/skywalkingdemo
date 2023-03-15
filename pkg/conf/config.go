package conf

import (
	"gopkg.in/yaml.v3"
	"os"
	"skywalkingdemo/pkg/utils"
	"strings"
)

type Config struct {
	Apps Apps `yaml:"Apps"`
}

type Item struct {
	Type     string `yaml:"Type"`
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Database string `yaml:"Database"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

type Apps struct {
	Database Item `yaml:"Database"`
}

func replaceEnvDefault(s string) string {
	// 解析环境变量的默认值
	parts := strings.SplitN(s, ":", 2)
	if len(parts) == 2 {
		// 获取环境变量的值，如果不存在则使用默认值
		key := strings.TrimLeft(parts[0], "${")
		value := strings.TrimRight(parts[1], "}")
		return utils.GetEnv(key, value)
	}
	return s
}

func ReadYamlConfig(path string) (*Config, error) {
	conf := &Config{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		err := yaml.NewDecoder(f).Decode(conf)
		if err != nil {
			return nil, err
		}
	}

	// 替换环境变量的默认值
	conf.Apps.Database.Host = replaceEnvDefault(conf.Apps.Database.Host)
	conf.Apps.Database.Database = replaceEnvDefault(conf.Apps.Database.Database)
	conf.Apps.Database.Username = replaceEnvDefault(conf.Apps.Database.Username)
	conf.Apps.Database.Password = replaceEnvDefault(conf.Apps.Database.Password)
	conf.Apps.Database.Port = replaceEnvDefault(conf.Apps.Database.Port)

	return conf, nil
}

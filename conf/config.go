package conf

import (
	"fmt"
	"irita-api/seelog"
	"strings"

	"github.com/spf13/viper"
)

// InitConfig 初始化配置
func InitConfig() error {
	v := viper.New()
	v.SetEnvPrefix(ConfigPrefix)
	v.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	v.SetEnvKeyReplacer(replacer)
	v.SetConfigFile(DefaultConfigFile)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("read config error, %v", err)
	}
	for _, key := range v.AllKeys() {
		seelog.Infof("%s=%v", key, v.Get(key))
	}
	Conf = &Config{}
	if err := v.Unmarshal(Conf); err != nil {
		return fmt.Errorf("unmarshal Config error, %v", err)
	}
	setLoggerLevel()
	return nil
}

func setLoggerLevel() {
	switch strings.ToUpper(Conf.LoggerConf.Level) {
	case "INFO":
		seelog.SetInfoLevel()
	case "WARN":
		seelog.SetWarnLevel()
	case "ERROR":
		seelog.SetErrorLevel()
	}
}

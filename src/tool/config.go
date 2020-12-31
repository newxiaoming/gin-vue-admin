package tool

import (
	"fmt"

	"github.com/spf13/viper"
)

type ruleConfig struct {
	name       string
	expression string
}

func GetRuleConfigFile() *viper.Viper {
	v := viper.New()
	var config ruleConfig
	// viper.SetConfigFile("./rule/rules.yaml")
	v.SetConfigName("rules")
	v.SetConfigType("yaml")
	v.AddConfigPath("./rule")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config)
	// fmt.Println(config.GetString("www_jiemianian_com.expression"))
	return v
}

func GetRuleConfig(key string, fileName string) string {
	v := viper.New()
	var config ruleConfig
	// viper.SetConfigFile("./rule/rules.yaml")
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./rule")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config)
	// fmt.Println(v.GetString("www_jiemianian_com.source"))
	fmt.Println(key)
	return v.GetString(key)

}

func GetConfig(key string) string {
	v := viper.New()
	var config ruleConfig
	// viper.SetConfigFile("./rule/rules.yaml")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	viper.Unmarshal(&config)
	return v.GetString(key)
}

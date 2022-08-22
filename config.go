package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

var MapConfig map[string]interface{}
var MapJson map[string]interface{}

func NewConfig(filename string, filepath string) (map[string]interface{}, error) {
	viper.AddConfigPath(filepath)
	file_info := strings.Split(filename, ".")
	viper.SetConfigName(file_info[0])
	viper.SetConfigType(file_info[1])
	viper.AutomaticEnv()

	if MapConfig == nil {
		if err := viper.ReadInConfig(); err != nil {
			return MapConfig, err
		}
		err := viper.Unmarshal(&MapConfig)
		if err != nil {
			return MapConfig, err
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		viper.AutomaticEnv()
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("read failed")
		}
		err := viper.Unmarshal(&MapConfig)
		if err != nil {
			fmt.Println("unmarshal failed")
		}
		fmt.Println("config updated")
	})
	viper.WatchConfig()

	return MapConfig, nil

}

func Common(key string, deep_key *string, default_val string) interface{} {
	value, found := MapConfig[key]

	if deep_key == nil {
		//str := fmt.Sprintf("%v", value)
		return value

	}
	if !found {
		return ""
	}
	b, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &MapJson)
	if err != nil {
		fmt.Println("error:", err)
	}
	deep_value, found := MapJson[*deep_key]
	if !found {
		return default_val
	}
	return deep_value
}

func GetKeyString(key string, deep_key *string, default_val string) string {
	p := Common(key, deep_key, default_val)
	return fmt.Sprintf("%v", p)
}

func GetKeyInt(key string, deep_key *string, default_value string) int {
	val := GetKeyString(key, deep_key, default_value)
	int_val, _ := strconv.Atoi(val)
	return int_val
}

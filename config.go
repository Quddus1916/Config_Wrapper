package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
)

type GetConfigParamAsString func(string, *string, string) string
type GetConfigParamAsInt64 func(string, *string, string) int64
type GetConfigParamAsFloat64 func(string, *string, string) float64
type Config struct {
	GetConfigParamAsString  GetConfigParamAsString
	GetConfigParamAsInt64   GetConfigParamAsInt64
	GetConfigParamAsFloat64 GetConfigParamAsFloat64
}

const bitSize = 64 // Don't think about it to much. It's just 64 bits.

var MapConfig map[string]interface{}
var MapJson map[string]interface{}

func Common(key string, deep_key *string, default_val string) interface{} {
	value, found := MapConfig[key]

	if !found {
		return default_val
	}

	if deep_key == nil {
		//str := fmt.Sprintf("%v", value)
		return value

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

func InitConfig(filepath string) (*Config, error) {
	var config = new(Config)
	viper.SetConfigFile(filepath)
	viper.AutomaticEnv()

	if MapConfig == nil {
		if err := viper.ReadInConfig(); err != nil {
			return config, err
		}
		err := viper.Unmarshal(&MapConfig)
		if err != nil {
			return config, err
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

	config = &Config{

		GetConfigParamAsString: func(key string, deep_key *string, default_val string) string {
			val := Common(key, deep_key, default_val)
			return fmt.Sprintf("%v", val)
		},
		GetConfigParamAsInt64: func(key string, deep_key *string, default_val string) int64 {
			val := Common(key, deep_key, default_val)
			Num, err := strconv.ParseInt(fmt.Sprintf("%v", val), 10, bitSize)
			if err != nil {
				fmt.Println("error:", err)
			}
			return Num
		},
		GetConfigParamAsFloat64: func(key string, deep_key *string, default_val string) float64 {
			val := Common(key, deep_key, default_val)
			Num, err := strconv.ParseFloat(fmt.Sprintf("%v", val), bitSize)
			if err != nil {
				fmt.Println("error:", err)
			}
			return Num
		},
	}

	return config, nil

}

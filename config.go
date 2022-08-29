package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"reflect"
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
type KeyFunc struct {
	Key          string
	CallBackFunc func()
}

var Pair []KeyFunc

const bitSize = 64 // Don't think about it to much. It's just 64 bits.
var MapConfig map[string]interface{}
var MapJson map[string]interface{}

func Common(key string, deep_key *string, default_val string) interface{} {

	value, found := MapConfig[key]
	if !found {
		return default_val
	}
	if deep_key == nil {
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

func MisMatchedKey(old map[string]interface{}, updated map[string]interface{}) string {
	for k, _ := range old {
		eq := reflect.DeepEqual(old[k], updated[k])
		if eq {

		} else {
			return k
		}
	}
	return ""
}

func CallFuncIfExists(key string) bool {
	for _, v := range Pair {
		if v.Key == key {
			v.CallBackFunc()
			return true
		}
	}
	return false
}

func InitConfig(filepath string, pair []KeyFunc) (*Config, error) {

	Pair = pair
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

		var OldConfig = make(map[string]interface{})
		fmt.Println("Config file changed:", e.Name)

		b, err := json.Marshal(MapConfig)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(b, &OldConfig)
		if err != nil {
			fmt.Println("error:", err)
		}

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("read failed")
		}

		err = viper.Unmarshal(&MapConfig)
		if err != nil {
			fmt.Println("unmarshal failed")
		}

		fmt.Println("config updated & Checking for any call_back func")

		key := MisMatchedKey(OldConfig, MapConfig)
		if key != "" {
			ok := CallFuncIfExists(key)
			if ok {
				fmt.Println("ok")
			}
		}

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
				Num, err := strconv.ParseInt(default_val, 10, bitSize)
				if err != nil {
					fmt.Println("error:", err)
				}
				return Num
			}
			return Num
		},
		GetConfigParamAsFloat64: func(key string, deep_key *string, default_val string) float64 {
			val := Common(key, deep_key, default_val)
			Num, err := strconv.ParseFloat(fmt.Sprintf("%v", val), bitSize)
			if err != nil {
				fmt.Println("error:", err)
				Num, err := strconv.ParseFloat(default_val, bitSize)
				if err != nil {
					fmt.Println("error:", err)
				}
				return Num
			}
			return Num
		},
	}

	return config, nil

}

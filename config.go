package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
)

type getConfigParamAsString func(string, *string, string) string
type getConfigParamAsInt64 func(string, *string, string) int64
type getConfigParamAsFloat64 func(string, *string, string) float64
type getParamAsStruct func(string, *string, string, any)

type Config struct {
	GetConfigParamAsString  getConfigParamAsString
	GetConfigParamAsInt64   getConfigParamAsInt64
	GetConfigParamAsFloat64 getConfigParamAsFloat64
	GetParamAsStruct        getParamAsStruct
}

var config *Config

type KeyFunc struct {
	Key          string
	CallBackFunc func()
}

var Pair []KeyFunc
var in_it bool

const bitSize = 64 // Don't think about it to much. It's just 64 bits.
var MapConfig map[string]interface{}
var MapJson map[string]interface{}
var OldConfig = make(map[string]interface{})

func GetConfig() *Config {
	if config == nil || MapConfig == nil || in_it == false {
		fmt.Println("plz.. initialize config first")
	}
	return config
}

func Decode(value interface{}) {
	b, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &MapJson)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func Common(key string, deep_key *string, default_val string) interface{} {
	var old_value interface{}
	value, found := MapConfig[key]
	if !found {
		old_value, found = OldConfig[key]
		if !found {
			return default_val
		}
		return old_value
	}
	if deep_key == nil {
		fmt.Println(&value)
		return value
	}

	Decode(value)

	deep_value, found := MapJson[*deep_key]
	if !found {
		Decode(old_value)
		deep_value_old, found := MapJson[*deep_key]
		if !found {

			return default_val
		}
		return deep_value_old
	}

	return deep_value
}

func MisMatchedKey(old map[string]interface{}, updated map[string]interface{}) []string {
	var keys []string
	for k, _ := range old {
		eq := reflect.DeepEqual(old[k], updated[k])
		if eq {

		} else {
			keys = append(keys, k)
		}
	}
	return keys
}

func CallFuncIfExists(key []string) bool {
	fmt.Println(key)
	for _, K := range key {
		for _, v := range Pair {
			if v.Key == K {
				v.CallBackFunc()

			}
		}
	}
	return true
}

func InitConfig(file_path string, pair []KeyFunc) error {
	if in_it == true {
		fmt.Println("already initialized")
		return nil
	}
	in_it = true
	_, err := os.Open(file_path)

	if err == nil {
		viper.SetConfigFile(file_path)
	} else {
		filename := filepath.Base(file_path)
		file := fmt.Sprintf("%s%s", "./", filename)
		viper.SetConfigFile(file)
	}

	Pair = pair
	viper.AutomaticEnv()

	if MapConfig == nil {
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		err := viper.Unmarshal(&MapConfig)
		if err != nil {
			return err
		}
	}

	viper.OnConfigChange(func(e fsnotify.Event) {

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
		if key != nil {
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
		GetParamAsStruct: func(key string, deep_key *string, default_val string, a any) {
			val := Common(key, deep_key, default_val)
			mapstructure.Decode(val, &a)
			fmt.Println(a)
			return
		},
	}

	return nil

}

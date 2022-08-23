package main

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"strconv"
)

const bitSize = 64 // Don't think about it to much. It's just 64 bits.

var MapConfig map[string]interface{}
var MapJson map[string]interface{}

func InitConfig(filepath string) (map[string]interface{}, error) {

	viper.SetConfigFile(filepath)
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

func GetConfigParamAsString(key string, deep_key *string, default_val string) string {
	val := Common(key, deep_key, default_val)
	return fmt.Sprintf("%v", val)
}

func GetConfigParamAsInt64(key string, deep_key *string, default_value string) int64 {
	val := Common(key, deep_key, default_value)
	Num, err := strconv.ParseInt(fmt.Sprintf("%v", val), 10, bitSize)
	if err != nil {
		fmt.Println("error:", err)
	}
	return Num
}

func GetConfigParamAsFloat64(key string, deep_key *string, default_value string) float64 {
	val := Common(key, deep_key, default_value)
	Num, err := strconv.ParseFloat(fmt.Sprintf("%v", val), bitSize)
	if err != nil {
		fmt.Println("error:", err)
	}
	return Num
}

func main() {
	_, err := InitConfig("./config.dev.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	val := GetConfigParamAsInt64("port", nil, "1010")
	fmt.Println(val)
	limit := "limit"
	port := "port"
	val2 := GetConfigParamAsInt64("app", &port, "1010")
	fmt.Println(val2)
	val3 := GetConfigParamAsFloat64("app", &limit, "1010")
	fmt.Println(val3)

}

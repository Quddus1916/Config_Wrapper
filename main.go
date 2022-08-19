package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	//"log"
	"strconv"
	"strings"
)

var MapConfig map[string]string
var MapJson map[string]string

func NewConfig(filename string, filepath string) (map[string]string, error) {
	viper.AddConfigPath(".")
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

func GetKeyLikeString(key string, default_val string) string {
	value, found := MapConfig[key]
	if !found {
		return default_val
	}
	return value
}

func GetKeyLikeInt(key string, default_value int) int {
	value, found := MapConfig[key]
	if !found {
		return default_value
	}
	int_val, _ := strconv.Atoi(value)
	return int_val
}

func GetKeyLikeJson(key string) map[string]string {
	value, found := MapConfig[key]
	if !found {
		return MapJson
	}
	err := json.Unmarshal([]byte(value), &MapJson)
	if err != nil {
		fmt.Println(err.Error())
	}
	return MapJson
}

// func main() {

// 	config, err := NewConfig("app.env", ".")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	val := GetKeyLikeString("port", "1010")
// 	fmt.Println(val)
// 	val2 := GetKeyLikeInt("port", 1010)
// 	fmt.Println(val2)
// 	val3 := GetKeyLikeJson("email")
// 	fmt.Print(val3)
// 	fmt.Println(config)
// }

package Config

import (
	//"fmt"
	"github.com/spf13/viper"
	"strings"
)

// type Config struct {
// 	Port         string `mapstructure:"PORT"`
// 	SecretKey    string `mapstructure:"SECRETKEY"`
// 	SqlUri       string `mapstructure:"SQL_URI"`
// 	SqlDb        string `mapstructure:"SQL_DB_NAME"`
// 	Username     string `mapstructure:"USER_NAME"`
// 	Email        string `mapstructure:"EMAIL"`
// 	SmtpHost     string `mapstructure:"SMTP_HOST"`
// 	SmtpPort     string `mapstructure:"SMTP_PORT"`
// 	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
// 	Link         string `mapstructure:"LINK"`
// }

func NewConfig(config_struct Config, filename string) (Config, error) {
	viper.AddConfigPath(".")
	file_info := strings.Split(filename, ".")
	viper.SetConfigName(file_info[0])
	viper.SetConfigType(file_info[1])
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return config_struct, err
	}

	err := viper.Unmarshal(&config_struct)
	if err != nil {
		return config_struct, err
	}

	return config_struct, err

}

// func main() {
// 	config, err := NewConfig(Config{}, "app.env")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	fmt.Println(config.Port)
// }

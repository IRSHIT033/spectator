package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv                 string `mapstructure:"APP_ENV"`
	ServerAddress          string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout         int    `mapstructure:"CONTEXT_TIMEOUT"`
	MongoURI               string `mapstructure:"MONGO_URI"`
	DBname                 string `mapstructure:"DB_NAME"`
	AccessTokenExpiryHour  int    `mapstructure:"ACCESS_TOKEN_EXPIRY_HOUR"`
	RefreshTokenExpiryHour int    `mapstructure:"REFRESH_TOKEN_EXPIRY_HOUR"`
	AccessTokenSecret      string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret     string `mapstructure:"REFRESH_TOKEN_SECRET"`
}

func InitConfig() *Config {
	cnfg := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&cnfg)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if cnfg.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &cnfg
}

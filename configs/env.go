package configs

import (
	"log"

	"github.com/spf13/viper"
)

var EnvConfigs *envConfigs

type envConfigs struct {
	MongoURI string `mapstructure:"MONGO_URI"`
	DBName   string `mapstructure:"DB_NAME"`
}

func InitEnvConfigs() {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&EnvConfigs); err != nil {
		log.Fatal(err)
	}
}

package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DbPort           int    `mapstructure:"DB_PORT"`
	DbHost           string `mapstructure:"DB_HOST"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`

	GrpcServerAddress string `mapstructure:"GRPC_SERVER_ADDRESS"`
	Origin            string `mapstructure:"CLIENT_ORIGIN"`
	Port              int    `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

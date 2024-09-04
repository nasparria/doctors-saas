// config/config.go
package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerPort    int      `mapstructure:"SERVER_PORT"`
	DBHost        string   `mapstructure:"DB_HOST"`
	DBPort        int      `mapstructure:"DB_PORT"`
	DBUser        string   `mapstructure:"DB_USER"`
	DBPassword    string   `mapstructure:"DB_PASSWORD"`
	DBName        string   `mapstructure:"DB_NAME"`
	KafkaBrokers  []string `mapstructure:"KAFKA_BROKERS"`
	KafkaGroupID  string   `mapstructure:"KAFKA_GROUP_ID"`
	EmailAPIToken string   `mapstructure:"EMAIL_API_TOKEN"`
	EmailFrom     string   `mapstructure:"EMAIL_FROM"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env") // Use .env file
	viper.AddConfigPath(".")    // Look for .env in the root directory

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

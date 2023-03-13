package cfg

import (
	"github.com/spf13/viper"
)

type Config struct {
	Grpc string `mapstructure:"GRPC_PORT"`
	Http string `mapstructure:"HTTP_PORT"`

	TokensAddr string `mapstructure:"TOKENS_GRPC_ADDR"`

	SmtpHost  string `mapstructure:"SMTP_HOST"`
	SmtpPort  string `mapstructure:"SMTP_PORT"`
	SmtpEmail string `mapstructure:"SMTP_EMAIL"`
	SmtpPass  string `mapstructure:"SMTP_PASS"`

	ApiHeader string `mapstructure:"API_HEADER"`
	ApiSecret string `mapstructure:"API_SECRET"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("example")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

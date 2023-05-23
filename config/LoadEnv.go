package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	Env            string `mapstructure:"ENVIRONMENT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_NAME"`
	DBUserPassword string `mapstructure:"POSTGRES_PASS"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ServerPort     string `mapstructure:"SERVER_PORT"`
	Timeout        int    `mapstructure:"TIMEOUT"`

	AccTokenPrivateKey string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccTokenPublicKey  string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	AccTokenExpireIn   string `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	AccTokenMaxEge     string `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	RefTokenExpireIn   string `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	RefTokenMaxAge     string `mapstructure:"REFRESH_TOKEN_MAXAGE"`

	EmailFrom    string `mapstructure:"EMAIL_FROM"`
	SmtpHost     string `mapstructure:"SMTP_HOST"`
	SmtpUser     string `mapstructure:"SMTP_USER"`
	SmtpPassword string `mapstructure:"SMTP_PASSWORD"`
	SmtpPort     int    `mapstructure:"SMTP_PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

var MyEnv Config

func init() {
	log.Println(" Start Load Config")

	var err error
	if MyEnv, err = LoadEnv("."); err != nil {
		log.Println("Load Config Failed", err.Error())
		panic(err)
	}
}

func LoadEnv(path string) (config Config, err error) {
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

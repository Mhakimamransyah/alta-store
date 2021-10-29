package config

import (
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type ConfigIPForwarding struct {
	Enabled bool   `mapstructure:"enabled"`
	IP      string `mapstructure:"ip"`
	Port    string `mapstructure:"port"`
}

//AppConfig Application configuration
type AppConfig struct {
	AppPort        int    `mapstructure:"app_port"`
	AppEnvironment string `mapstructure:"app_environment"`
	DbDriver       string `mapstructure:"db_driver"`
	DbAddress      string `mapstructure:"db_address"`
	DbPort         int    `mapstructure:"db_port"`
	DbUsername     string `mapstructure:"db_username"`
	DbPassword     string `mapstructure:"db_password"`
	DbName         string `mapstructure:"db_name"`
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

//GetConfig Initiatilize config in singleton way
func GetConfig() *AppConfig {
	if appConfig != nil {
		return appConfig
	}

	lock.Lock()
	defer lock.Unlock()

	//re-check after locking
	if appConfig != nil {
		return appConfig
	}

	appConfig = initConfig()

	return appConfig
}

func initConfig() *AppConfig {
	godotenv.Load(".env")

	var defaultConfig AppConfig
	var finalConfig AppConfig

	defaultConfig.AppPort = 8000
	defaultConfig.AppEnvironment = ""
	defaultConfig.DbDriver = "mysql"
	defaultConfig.DbAddress = "127.0.0.1"
	defaultConfig.DbPort = 3306
	defaultConfig.DbUsername = "your DB username"
	defaultConfig.DbPassword = "your DB password"
	defaultConfig.DbName = "your DB name"

	//use this if .env file (dont forget to run "source PATH_TO/.env" example "source config/.env")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("altastore")
	viper.BindEnv("app_port")
	viper.BindEnv("app_environment")
	viper.BindEnv("db_driver")
	viper.BindEnv("db_address")
	viper.BindEnv("db_port")
	viper.BindEnv("db_username")
	viper.BindEnv("db_password")
	viper.BindEnv("db_name")
	// viper.ReadInConfig()

	err := viper.Unmarshal(&finalConfig)
	if err != nil {
		log.Info("failed to extract config, will use default value")
		return &defaultConfig
	}

	return &finalConfig
}

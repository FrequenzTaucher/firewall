package services

import (
	"fmt"
	"spamtrawler/helper"
	"spamtrawler/repository"
	"spamtrawler/repository/models"

	"github.com/fsnotify/fsnotify"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/spf13/viper"

	"github.com/labstack/gommon/log"
)

var DB *mongo.Database
var Configuration models.Configuration
var RootDirectory string

func init() {
	RootDirectory = helper.GetRootDirectory()
	Configuration = readInConfig()
	DB = repository.GetDbConnection(Configuration)
}

func readInConfig() models.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(RootDirectory + "/files/")

	viper.WatchConfig()

	/*
	 * ToDo: Reload Config on change
	 */
	viper.OnConfigChange(func(e fsnotify.Event) {
		readInConfig()
		fmt.Println("Config file changed: ", e.Name)
	})

	var configuration models.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading  config file: %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	return configuration
}

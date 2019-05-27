package app

import (
	"fmt"
	"log"
	"os"
	"spamtrawler/app/models"
	"spamtrawler/app/repository"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var MongoDB *mongo.Database
var Configuration models.Configuration

func init() {
	Configuration = ReadInConfig()

	fmt.Println(Configuration)

	MongoDB = repository.MongoGetDbConnection(&Configuration)

	fmt.Println(MongoDB)
}

func Start() {
	fmt.Println("App Started!")
}

func GetRootDirectory() string {
	dir, err := os.Getwd()
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func ReadInConfig() models.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(GetRootDirectory() + "/files/")

	viper.WatchConfig()

	/*
	 * ToDo: Reload Config on change
	 */
	viper.OnConfigChange(func(e fsnotify.Event) {
		ReadInConfig()
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

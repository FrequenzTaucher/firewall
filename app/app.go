package app

import (
	"spamtrawler/app/helper"
	"spamtrawler/app/models"
	"spamtrawler/app/repository"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var DB *mongo.Database
var Configuration models.Configuration
var RootDirectory string

func init() {
	RootDirectory = helper.GetRootDirectory()
	Configuration = helper.ReadInConfig()
	DB = repository.GetDbConnection(Configuration)
}

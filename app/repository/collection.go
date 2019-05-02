package repository

import (
	"context"
	"net/http"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func CreateCollectionItem(collectionName string, d bson.D) (result *mongo.InsertOneResult, err error) {

	collection := MongoDB.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err = collection.InsertOne(ctx, d)

	return result, nil

}

func GetAllCollectionItems(collectionName string, c echo.Context) ([]bson.M, error) {
	var asns []bson.M

	collection := MongoDB.Collection(collectionName)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var asn bson.M
		_ = cursor.Decode(&asn)
		asns = append(asns, asn)
	}
	if err := cursor.Err(); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return nil, err
	}
	return asns, err
}

func GetCollectionItemById(collectionName string, c echo.Context) (result bson.M, err error) {

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	collection := MongoDB.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return nil, err
	}
	return result, nil
}

func DeleteCollectionItemById(collectionName string, c echo.Context) (result bson.M, err error) {

	id, err := primitive.ObjectIDFromHex(c.Param("id"))

	collection := MongoDB.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = collection.FindOneAndDelete(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return nil, err
	}
	return result, nil
}

func UpdateCollectionItemById(collectionName string, c echo.Context, data bson.D) (result *mongo.UpdateResult, err error) {
	collection := MongoDB.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err = collection.UpdateOne(
		ctx,
		bson.D{
			{"_id", c.Param("id")},
		},
		bson.D{
			{"$set", data},
		},
	)

	return result, err
}

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

type Country struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ISO    string             `json:"iso,omitempty" bson:"iso,omitempty"`
	STATUS string             `json:"status,omitempty" bson:"status,omitempty"`
}

func CreateCountry(c echo.Context) (result *mongo.InsertOneResult, err error) {
	d := new(Country)

	if err = c.Bind(d); err != nil {
		return nil, err
	}

	collection := MongoDB.Collection("countries")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err = collection.InsertOne(ctx, d)

	return result, nil
}

func GetAllCountries(c echo.Context) ([]Country, error) {
	var countries []Country
	collection := MongoDB.Collection("countries")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var country Country
		_ = cursor.Decode(&country)
		countries = append(countries, country)
	}
	if err := cursor.Err(); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return nil, err
	}
	return countries, err
}

func GetCountry(c echo.Context) (country Country, err error) {
	d := new(Country)

	if err = c.Bind(d); err != nil {
		return
	}

	collection := MongoDB.Collection("countries")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = collection.FindOne(ctx, Country{ID: d.ID}).Decode(&country)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return country, err
	}
	return country, nil
}

func DeleteCountry(c echo.Context) (country Country, err error) {
	d := new(Country)

	if err = c.Bind(d); err != nil {
		return
	}

	collection := MongoDB.Collection("countries")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = collection.FindOneAndDelete(ctx, Country{ID: d.ID}).Decode(&country)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return country, err
	}
	return country, err
}

func UpdateCountry(c echo.Context) (result *mongo.UpdateResult, err error) {
	d := new(Country)

	if err = c.Bind(d); err != nil {
		return nil, err
	}

	collection := MongoDB.Collection("countries")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err = collection.UpdateOne(
		ctx,
		bson.D{
			{"_id", d.ID},
		},
		bson.D{
			{"$set", bson.D{
				{"iso", d.ISO},
				{"status", d.STATUS},
			},
			},
		},
	)

	return result, err
}

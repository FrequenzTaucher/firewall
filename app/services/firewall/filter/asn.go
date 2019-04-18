package services

import (
	"context"
	"encoding/json"
	"net/http"
	"spamtrawler/app"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type ASN struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ASN    string             `json:"asn,omitempty" bson:"asn,omitempty"`
	STATUS string             `json:"status,omitempty" bson:"status,omitempty"`
}

func FilterAsn(asn uint) bool {
	asnList := make(map[uint]struct{})
	asnList[432489] = struct{}{}

	//ping := db.Ping(context.TODO(), nil)

	//fmt.Println(ping)

	_, found := asnList[asn]

	return found
}

func CreateAsn(c echo.Context) (err error) {
	d := new(ASN)

	if err = c.Bind(d); err != nil {
		return
	}

	collection := app.DB.Collection("asn")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, d)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func GetAllAsn(c echo.Context) error {
	var asns []ASN
	collection := app.DB.Collection("asn")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var asn ASN
		_ = cursor.Decode(&asn)
		asns = append(asns, asn)
	}
	if err := cursor.Err(); err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return err
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(asns)
}

func GetAsn(c echo.Context) (err error) {
	d := new(ASN)

	if err = c.Bind(d); err != nil {
		return
	}

	var asn ASN
	collection := app.DB.Collection("asn")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = collection.FindOne(ctx, ASN{ID: d.ID}).Decode(&asn)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(asn)
}

func DeleteAsn(c echo.Context) (err error) {
	d := new(ASN)

	if err = c.Bind(d); err != nil {
		return
	}

	var asn ASN
	collection := app.DB.Collection("asn")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err = collection.FindOneAndDelete(ctx, ASN{ID: d.ID}).Decode(&asn)
	if err != nil {
		c.Response().WriteHeader(http.StatusInternalServerError)
		_, _ = c.Response().Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(asn)
}

func UpdateAsn(c echo.Context) (err error) {
	d := new(ASN)

	if err = c.Bind(d); err != nil {
		return
	}

	collection := app.DB.Collection("asn")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.UpdateOne(
		ctx,
		bson.D{
			{"_id", d.ID},
		},
		bson.D{
			{"$set", bson.D{
				{"asn", d.ASN},
				{"status", d.STATUS},
			},
			},
		},
	)

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

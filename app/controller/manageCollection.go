package controller

import (
	"encoding/json"
	"net/http"
	"spamtrawler/app/models"
	"spamtrawler/app/repository"

	"github.com/mongodb/mongo-go-driver/bson"

	"github.com/labstack/echo"
)

/*func FilterAsn(asn uint) bool {
	asnList := make(map[uint]struct{})
	asnList[432489] = struct{}{}

	//ping := db.Ping(context.TODO(), nil)

	//fmt.Println(ping)

	_, found := asnList[asn]

	return found
}*/

func CreateCollectionItem(c echo.Context) (err error) {

	d := new(models.ASN)

	if err = c.Bind(d); err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	data := bson.D{
		{"asn", d.ASN},
		{"status", d.STATUS},
	}

	result, err := repository.CreateCollectionItem(c.Param("collection"), data)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func GetAllCollectionItems(c echo.Context) error {
	result, err := repository.GetAllCollectionItems(c.Param("collection"), c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func GetCollectionItemById(c echo.Context) (err error) {

	result, err := repository.GetCollectionItemById(c.Param("collection"), c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func DeleteCollectionItemById(c echo.Context) (err error) {
	result, err := repository.DeleteCollectionItemById(c.Param("collection"), c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func UpdateCollectionItemById(c echo.Context) (err error) {
	d := new(models.ASN)

	if err = c.Bind(d); err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	data := bson.D{
		{"asn", d.ASN},
		{"status", d.STATUS},
	}

	result, err := repository.UpdateCollectionItemById(c.Param("collection"), d.ID, data)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

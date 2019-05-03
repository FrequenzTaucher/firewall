package controller

import (
	"encoding/json"
	"net/http"
	"spamtrawler/app/repository"
	"spamtrawler/app/services"

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

func CreateFirewallListItem(c echo.Context) (err error) {

	data, err := services.PrepareGenericFirewallListDataForCollection(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result, err := repository.CreateCollectionItem(c.Param("collection"), data)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func GetAllFirewallListItems(c echo.Context) error {

	result, err := repository.GetAllCollectionItems(c.Param("collection"), c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	data := map[string][]bson.M{"data": result}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(data)
}

func GetFirewallListItemById(c echo.Context) (err error) {

	result, err := repository.GetCollectionItemById(c.Param("collection"), c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func DeleteFirewallListItemById(c echo.Context) (err error) {

	result, err := repository.DeleteCollectionItemById(c.Param("collection"), c)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func UpdateFirewallListItemById(c echo.Context) (err error) {

	data, err := services.PrepareGenericFirewallListDataForCollection(c)

	result, err := repository.UpdateCollectionItemById(c.Param("collection"), c, data)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}
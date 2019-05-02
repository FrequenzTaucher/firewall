package services

import (
	"net/http"
	"spamtrawler/app/models"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson"
)

func PrepareGenericFirewallListDataForCollection(c echo.Context) (data bson.D, err error) {
	d := new(models.GenericFirewallItem)

	if err = c.Bind(d); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return bson.D{
		{"value", d.VALUE},
		{"status", d.STATUS},
	}, nil
}

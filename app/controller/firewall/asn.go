package firewall

import (
	"encoding/json"
	"net/http"
	"spamtrawler/app/repository"

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
	result, err := repository.CreateAsn(c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func GetAllAsn(c echo.Context) error {
	result, err := repository.GetAllAsn(c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func GetAsn(c echo.Context) (err error) {
	result, err := repository.GetAsn(c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func DeleteAsn(c echo.Context) (err error) {
	result, err := repository.DeleteAsn(c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

func UpdateAsn(c echo.Context) (err error) {
	result, err := repository.UpdateAsn(c)

	if err != nil {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		c.Response().WriteHeader(http.StatusInternalServerError)
		return json.NewEncoder(c.Response()).Encode(err)
	}

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(result)
}

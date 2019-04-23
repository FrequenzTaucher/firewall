package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spamtrawler/app"
	"spamtrawler/app/controller/firewall"
	"spamtrawler/app/models"
	geoip "spamtrawler/app/services/geoip"
	machine "spamtrawler/app/services/machine"

	"github.com/spf13/viper"

	"github.com/labstack/echo"
)

func RouteHandler(e *echo.Echo) {
	fmt.Println(viper.AllKeys())
	// Route => handler
	e.POST("/filter", func(c echo.Context) (err error) {
		d := new(models.VisitorDetailsQuery)

		if err = c.Bind(d); err != nil {
			return
		}

		if d.ApiKey == app.Configuration.Auth.ApiKey {
			fmt.Println(app.Configuration.Environment.Mode)

			//e.Logger.Debug(d)
			visitorGeoData := geoip.GetGeoDataFromIp(app.RootDirectory, d.IpAddress)
			visitorNetworkData := geoip.GetNetworkDataFromIp(app.RootDirectory, d.IpAddress)
			machineData := machine.GetMachineData(d.UserAgent)

			blockedFlag := false
			blockedFlag = firewall.FilterAsn(visitorNetworkData.AutonomousSystemNumber)
			if blockedFlag == false {
				blockedFlag = firewall.FilterCountry(visitorGeoData.CountryIsoCode)
			}

			data := models.FirewallOutputResult{
				QueryData:   *d,
				GeoData:     *visitorGeoData,
				NetworkData: *visitorNetworkData,
				MachineData: *machineData,
				Blocked:     blockedFlag,
			}

			c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
			c.Response().WriteHeader(http.StatusOK)
			return json.NewEncoder(c.Response()).Encode(data)
		}

		// For invalid credentials
		return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid API credentials")

		//return c.JSON(http.StatusOK, data)
	})

	asn := e.Group("/asn")
	asn.POST("/create", func(c echo.Context) (err error) {
		return firewall.CreateAsn(c)
	})
	asn.GET("/get/:id", func(c echo.Context) (err error) {
		return firewall.GetAsn(c)
	})
	asn.GET("/all", func(c echo.Context) (err error) {
		return firewall.GetAllAsn(c)
	})
	asn.POST("/update", func(c echo.Context) (err error) {
		return firewall.UpdateAsn(c)
	})
	asn.GET("/delete/:id", func(c echo.Context) (err error) {
		return firewall.DeleteAsn(c)
	})

	country := e.Group("/country")
	country.POST("/create", func(c echo.Context) (err error) {
		return firewall.CreateCountry(c)
	})
	country.GET("/get/:id", func(c echo.Context) (err error) {
		return firewall.GetCountry(c)
	})
	country.GET("/all", func(c echo.Context) (err error) {
		return firewall.GetAllCountry(c)
	})
	country.POST("/update", func(c echo.Context) (err error) {
		return firewall.UpdateCountry(c)
	})
	country.GET("/delete/:id", func(c echo.Context) (err error) {
		return firewall.DeleteCountry(c)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1232"))
}

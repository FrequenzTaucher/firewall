package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spamtrawler/app/services"
	firewall "spamtrawler/app/services/firewall/filter"
	geoip "spamtrawler/app/services/geoip"
	machine "spamtrawler/app/services/machine"
	"spamtrawler/repository/models"

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

		if d.ApiKey == services.Configuration.Auth.ApiKey {
			fmt.Println(services.Configuration.Environment.Mode)

			//e.Logger.Debug(d)
			visitorGeoData := geoip.GetGeoDataFromIp(services.RootDirectory, d.IpAddress)
			visitorNetworkData := geoip.GetNetworkDataFromIp(services.RootDirectory, d.IpAddress)
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

	e.GET("/getAllAsn", func(c echo.Context) (err error) {
		return firewall.GetAllAsn(c)
	})
	e.GET("/getAsn/:id", func(c echo.Context) (err error) {
		return firewall.GetAsn(c)
	})
	e.GET("/deleteAsn/:id", func(c echo.Context) (err error) {
		return firewall.DeleteAsn(c)
	})
	e.POST("/createAsn", func(c echo.Context) (err error) {
		return firewall.CreateAsn(c)
	})
	e.POST("/updateAsn", func(c echo.Context) (err error) {
		return firewall.UpdateAsn(c)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1232"))
}

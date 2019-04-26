package routes

import (
	"spamtrawler/app/controller"

	"github.com/labstack/echo"
)

func RouteHandler(e *echo.Echo) {
	//fmt.Println(viper.AllKeys())
	// Route => handler

	/*
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
				blockedFlag = controller.FilterAsn(visitorNetworkData.AutonomousSystemNumber)
				if blockedFlag == false {
					blockedFlag = controller.FilterCountry(visitorGeoData.CountryIsoCode)
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
		})*/

	asn := e.Group("/manage")
	asn.POST("/:collection/create", func(c echo.Context) (err error) {
		return controller.CreateCollectionItem(c)
	})
	asn.GET("/:collection/get/:id", func(c echo.Context) (err error) {
		return controller.GetCollectionItemById(c)
	})
	asn.GET("/:collection/all", func(c echo.Context) (err error) {
		return controller.GetAllCollectionItems(c)
	})
	asn.PUT("/:collection/update", func(c echo.Context) (err error) {
		return controller.UpdateCollectionItemById(c)
	})
	asn.DELETE("/:collection/delete/:id", func(c echo.Context) (err error) {
		return controller.DeleteCollectionItemById(c)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1232"))
}

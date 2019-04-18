package models

import "github.com/avct/uasurfer"

type VisitorDetailsQuery struct {
	ApiKey       string `json:"apiKey" form:"apiKey" query:"apiKey"`
	IpAddress    string `json:"ipAddress" form:"ipAddress" query:"ipAddress"`
	UserName     string `json:"userName" form:"userName" query:"userName"`
	EmailAddress string `json:"emailAddress" form:"emailAddress" query:"emailAddress"`
	UserAgent    string `json:"userAgent" form:"userAgent" query:"userAgent"`
}

type VisitorGeoData struct {
	CountryIsoCode    string  `json:"countryIsoCode"`
	CountryName       string  `json:"countryName"`
	ContinentIsoCode  string  `json:"continentIsoCode"`
	ContinentName     string  `json:"continentName"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	IsInEuropeanUnion bool    `json:"isInEuropeanUnion"`
}

type VisitorNetworkData struct {
	AutonomousSystemNumber       uint   `json:"autonomousSystemNumber"`
	AutonomousSystemOrganization string `json:"autonomousSystemOrganization"`
	HostName                     string `json:"hostName"`
}

type MachineData struct {
	BrowserName    string           `json:"browserName"`
	BrowserVersion uasurfer.Version `json:"browserVersion"`
	OsPlatform     string           `json:"osPlatform"`
	OsName         string           `json:"osName"`
	OsVersion      uasurfer.Version `json:"osVersion"`
	DeviceType     string           `json:"deviceType"`
	IsBot          bool             `json:"isBot"`
}

type FirewallOutputResult struct {
	QueryData   VisitorDetailsQuery `json:"queryData"`
	GeoData     VisitorGeoData      `json:"geoData"`
	NetworkData VisitorNetworkData  `json:"networkData"`
	MachineData MachineData         `json:"machineData"`
	Blocked     bool                `json:"blocked"`
}

type FirewallDataOutputNew struct {
	QueryData                    VisitorDetailsQuery `json:"queryData"`
	CountryIsoCode               string              `json:"countryIsoCode"`
	CountryName                  string              `json:"countryName"`
	ContinentIsoCode             string              `json:"continentIsoCode"`
	ContinentName                string              `json:"continentName"`
	Latitude                     float64             `json:"latitude"`
	Longitude                    float64             `json:"longitude"`
	IsInEuropeanUnion            bool                `json:"isInEuropeanUnion"`
	AutonomousSystemNumber       uint                `json:"autonomousSystemNumber"`
	AutonomousSystemOrganization string              `json:"autonomousSystemOrganization"`
	HostName                     string              `json:"hostName"`
	BrowserName                  string              `json:"browserName"`
	BrowserVersion               uasurfer.Version    `json:"browserVersion"`
	OsPlatform                   string              `json:"osPlatform"`
	OsName                       string              `json:"osName"`
	OsVersion                    uasurfer.Version    `json:"osVersion"`
	DeviceType                   string              `json:"deviceType"`
	IsBot                        bool                `json:"isBot"`
}

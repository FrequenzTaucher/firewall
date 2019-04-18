package services

import (
	"log"
	"net"
	"spamtrawler/app/models"

	"github.com/oschwald/geoip2-golang"
)

func GetGeoDataFromIp(dir, ip string) *models.VisitorGeoData {
	db, err := geoip2.Open(dir + "/files/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ipAddress := net.ParseIP(ip)

	if ipAddress == nil {
		log.Fatal(err)
	}

	record, err := db.City(ipAddress)
	if err != nil {
		log.Fatal(err)
	}

	data := models.VisitorGeoData{
		CountryIsoCode:    record.Country.IsoCode,
		CountryName:       record.Country.Names["en"],
		ContinentIsoCode:  record.Continent.Code,
		ContinentName:     record.Continent.Names["en"],
		Latitude:          record.Location.Latitude,
		Longitude:         record.Location.Longitude,
		IsInEuropeanUnion: record.Country.IsInEuropeanUnion,
	}

	return &data
}

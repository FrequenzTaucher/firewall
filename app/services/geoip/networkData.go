package services

import (
	"log"
	"net"
	"spamtrawler/app/models"

	"github.com/oschwald/geoip2-golang"
)

func GetNetworkDataFromIp(dir, ip string) *models.VisitorNetworkData {
	db, err := geoip2.Open(dir + "/files/GeoLite2-ASN.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ipAddress := net.ParseIP(ip)

	if ipAddress == nil {
		log.Fatal(err)
	}

	record, err := db.ASN(ipAddress)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(record)
	data := models.VisitorNetworkData{
		AutonomousSystemNumber:       record.AutonomousSystemNumber,
		AutonomousSystemOrganization: record.AutonomousSystemOrganization,
		HostName:                     getNetworkHostName(ip),
	}

	return &data

}

func getNetworkHostName(ip string) string {

	host, err := net.LookupAddr(ip)

	if err == nil && len(host) > 0 {
		return host[0]
	}

	return ""
}

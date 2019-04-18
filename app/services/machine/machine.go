package services

import (
	"spamtrawler/app/models"
	"strings"

	"github.com/avct/uasurfer"
)

func GetMachineData(userAgent string) *models.MachineData {
	ua := uasurfer.Parse(userAgent)

	data := models.MachineData{
		BrowserName:    strings.Replace(ua.Browser.Name.String(), "Browser", "", 1),
		BrowserVersion: ua.Browser.Version,
		OsPlatform:     strings.Replace(ua.OS.Platform.String(), "Platform", "", 1),
		OsName:         strings.Replace(ua.OS.Name.String(), "OS", "", 1),
		OsVersion:      ua.OS.Version,
		DeviceType:     strings.Replace(ua.DeviceType.String(), "Device", "", 1),
		IsBot:          ua.IsBot(),
	}

	return &data

}

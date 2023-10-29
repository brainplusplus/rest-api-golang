package utils

import (
	"simple-ecommerce/types"
)

//var (
//	timezoneLoc *time.Location
//)
//
//func InitTimeZoneLocation() {
//	timezoneLocString := configs.GetConfigString("server.timezone")
//	SetTimeZoneLocation(timezoneLocString)
//}
//
//func GetTimeZoneLocation() *time.Location {
//	return timezoneLoc
//}
//
//func SetTimeZoneLocation(timezoneLocString string) {
//	timezoneLoc, _ = time.LoadLocation(timezoneLocString)
//}

func JsonDateToTimeString(jsonDate types.JsonDate) (resDate string) {
	return jsonDate.ToTimeString()
}

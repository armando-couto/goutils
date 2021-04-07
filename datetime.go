package goutils

import (
	"strings"
	"time"
)

const LAYOUT_DDMMYYYY = "02012006"
const LAYOUT_DD_MM_YYYY = "02/01/2006"
const LAYOUT_DD_MM_YYYY_HH_MM_SS = "02/01/2006 15:04:05"
const LAYOUT_YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
const LAYOUT_YYYYMMDD = "20060102"
const LAYOUT_YYMMDDHHMMSS = "060102150405"
const LAYOUT_YYYYMMDDHHMMSS = "20060102150405"
const LAYOUT_YYYYMM = "200601"
const LAYOUT_YYYY_MM_DD = "2006-01-02"
const LAYOUT_YYMMDD = "060102"
const LAYOUT_HHMMSS = "150405"
const LAYOUT_HH = "15"
const LAYOUT_HH_MM_SS = "15:04:05"
const LAYOUT_YYYY_MM_DDTHH_MM_SS_000Z = "2006-01-02T15:04:05.000Z"
const LAYOUT_YYYY_MM_DDTHH_MM_SS_000 = "2006-01-02 15:04:05"

var LOC, _ = time.LoadLocation("America/Fortaleza")

/*
	ConvertStringToTimeLayoutDDMMYYYY
*/
func ConvertStringToTimeLayoutDDMMYYYY(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_DDMMYYYY, value, LOC)
	return t
}

/*
	ConvertTimeToStringLayoutDD_MM_YYYY
*/
func ConvertTimeToStringLayoutDD_MM_YYYY(date time.Time) string {
	return date.Format(LAYOUT_DD_MM_YYYY)
}

/*
	ConvertStringToTimeLayoutDD_MM_YYYY
*/
func ConvertStringToTimeLayoutDD_MM_YYYY(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_DD_MM_YYYY, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutDD_MM_YYYY_HH_MM_SS
*/
func ConvertStringToTimeLayoutDD_MM_YYYY_HH_MM_SS(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_DD_MM_YYYY_HH_MM_SS, value, LOC)
	return t
}

/*
	ConvertStringDD_MM_YYYY
*/
func ConvertStringDD_MM_YYYY(date string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_DD_MM_YYYY, date, LOC)
	return t
}

/*
	ConverTimeToStrinLayoutYYYY_MM_DD
*/
func ConverTimeToStrinLayoutYYYY_MM_DD(data time.Time) string {
	data.Format(LAYOUT_YYYY_MM_DD)
	return strings.Split(data.String(), " ")[0]
}

/*
	ConvertStringToTimeLayoutYYYYMMDD
*/
func ConvertStringToTimeLayoutYYYYMMDD(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_YYYYMMDD, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutYYYY_MM_DD
*/
func ConvertStringToTimeLayoutYYYY_MM_DD(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_YYYY_MM_DD, value, LOC)
	return t
}

/*
	ConverTimeToStrinLayoutYYYYMMDD
*/
func ConverTimeToStrinLayoutYYYYMMDD(data time.Time) string {
	data.Format(LAYOUT_YYYYMMDD)
	return strings.ReplaceAll(strings.Split(data.String(), " ")[0], "-", "")
}

/*
	ConvertStringToTimeLayoutYYMMDDHHMMSS
*/
func ConvertStringToTimeLayoutYYMMDDHHMMSS(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_YYMMDDHHMMSS, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutYYYYMMDDHHMMSS
*/
func ConvertStringToTimeLayoutYYYYMMDDHHMMSS(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_YYYYMMDDHHMMSS, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayout_YYYY_MM_DD_HH_MM_SS o antigo nome era: ConvertStringToTimeLAYOUT_YYYY_MM_DD_HH_MM_SS
*/
func ConvertStringToTimeLayout_YYYY_MM_DD_HH_MM_SS(date time.Time) string {
	return date.Format(LAYOUT_YYYY_MM_DD_HH_MM_SS)
}

/*
	ConvertStringToTimeLayout_YYYY_MM_DD o antigo nome era: ConvertStringToTimeLAYOUT_YYYY_MM_DD
*/
func ConvertStringToTimeLayout_YYYY_MM_DD(date time.Time) string {
	return date.Format(LAYOUT_YYYY_MM_DD)
}

/*
	ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000Z
*/
func ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000Z(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_YYYY_MM_DDTHH_MM_SS_000Z, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000
*/
func ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_YYYY_MM_DDTHH_MM_SS_000, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutHHMMSS
*/
func ConvertStringToTimeLayoutHHMMSS(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_HHMMSS, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutHH_MM_SS
*/
func ConvertStringToTimeLayoutHH_MM_SS(value string) time.Time {
	t, _ := time.ParseInLocation(LAYOUT_HH_MM_SS, value, LOC)
	return t
}

/*
	ConvertStringToTimeLayoutDDMMYYYYHHMMSS
*/
func ConvertStringToTimeLayoutDDMMYYYYHHMMSS(d time.Time, h time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), h.Hour(), h.Minute(), h.Second(), 0, time.UTC)
}

/*
	RangeDate returns a date range function over start date to end date inclusive.
	After the end of the range, the range function returns a zero date, date.IsZero() is true.
*/
func RangeDate(end, start time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

/*
	DatePlusTime
*/
func DatePlusTime(date, timeOfDay time.Time) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", date.Format("2006-01-02")+" "+timeOfDay.Format("15:04:05"), LOC)
}

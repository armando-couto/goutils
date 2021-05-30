package goutils

import (
	"strings"
	"time"
)

const LAYOUT_DDMMYYYY = "02012006"
const LAYOUT_DD_MM_YYYY = "02/01/2006"
const LAYOUT_DD_MM_YYYY_HH_MM_SS = "02/01/2006 15:04:05"
const LAYOUT_YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
const LAYOUT_MM_DD_YYYY_HH_MM = "01-02-2006 15:04"
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

/*
	ConvertStringToTimeLayoutDDMMYYYY
*/
func ConvertStringToTimeLayoutDDMMYYYY(value string) time.Time {
	t, _ := time.Parse(LAYOUT_DDMMYYYY, value)
	return t
}

/*
	ConvertTimeToStringLayoutDD_MM_YYYY
*/
func ConvertTimeToStringLayoutDD_MM_YYYY(date time.Time) string {
	return date.Format(LAYOUT_DD_MM_YYYY)
}

func ConvertTimeToStringLayoutMM_DD_YYYY_HH_MM(date time.Time) string {
	return date.Format(LAYOUT_MM_DD_YYYY_HH_MM)
}

/*
	ConvertStringToTimeLayoutDD_MM_YYYY
*/
func ConvertStringToTimeLayoutDD_MM_YYYY(value string) time.Time {
	t, _ := time.Parse(LAYOUT_DD_MM_YYYY, value)
	return t
}

/*
	ConvertStringToTimeLayoutDD_MM_YYYY_HH_MM_SS
*/
func ConvertStringToTimeLayoutDD_MM_YYYY_HH_MM_SS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_DD_MM_YYYY_HH_MM_SS, value)
	return t
}

/*
	ConvertStringDD_MM_YYYY
*/
func ConvertStringDD_MM_YYYY(date string) time.Time {
	t, _ := time.Parse(LAYOUT_DD_MM_YYYY, date)
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
	t, _ := time.Parse(LAYOUT_YYYYMMDD, value)
	return t
}

/*
	ConvertStringToTimeLayoutYYYY_MM_DD
*/
func ConvertStringToTimeLayoutYYYY_MM_DD(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYY_MM_DD, value)
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
	t, _ := time.Parse(LAYOUT_YYMMDDHHMMSS, value)
	return t
}

/*
	ConvertStringToTimeLayoutYYYYMMDDHHMMSS
*/
func ConvertStringToTimeLayoutYYYYMMDDHHMMSS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYYMMDDHHMMSS, value)
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
	t, _ := time.Parse(LAYOUT_YYYY_MM_DDTHH_MM_SS_000Z, value)
	return t
}

/*
	ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000
*/
func ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYY_MM_DDTHH_MM_SS_000, value)
	return t
}

/*
	ConvertStringToTimeLayoutHHMMSS
*/
func ConvertStringToTimeLayoutHHMMSS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_HHMMSS, value)
	return t
}

/*
	ConvertStringToTimeLayoutHH_MM_SS
*/
func ConvertStringToTimeLayoutHH_MM_SS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_HH_MM_SS, value)
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
	return time.Parse("2006-01-02 15:04:05", date.Format("2006-01-02")+" "+timeOfDay.Format("15:04:05"))
}

/*
	Qual é o dia da Segunda-feira da data que passou
*/
func WeekStartDate(date time.Time) time.Time {
	if int(time.Sunday) == 0 { // Caso seja Domingo jogar para segunda
		date = date.Add(time.Hour * 24)
	}

	offset := (int(time.Monday) - int(date.Weekday()) - 7) % 7
	result := date.Add(time.Duration(offset*24) * time.Hour)
	return result
}

/*
	Qual é o dia da Sexta-feira da data que passou
*/
func WeekEndDate(date time.Time) time.Time {
	offset := int(time.Friday) - int(date.Weekday())
	result := date.Add(time.Duration(offset*24) * time.Hour)
	return result
}

package goutils

import (
	"fmt"
	"strconv"
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
const LAYOUT_HH_MM_SS = "15:04:05"
const LAYOUT_YYYY_MM_DDTHH_MM_SS_000Z = "2006-01-02T15:04:05.000Z"
const LAYOUT_YYYY_MM_DDTHH_MM_SS_000 = "2006-01-02 15:04:05"

func ConvertStringToInt(value string) int {
	i, _ := strconv.Atoi(value)
	return i
}

func ConvertIntToString(value int) string {
	s := strconv.Itoa(value)
	return s
}

func ConvertStringToFloatScale2Virgula(value string) float64 {
	value = strings.Replace(value, "%", "", 1)
	value = strings.Replace(value, ".", "", 1)
	value = strings.Replace(value, "R", "", 1)
	value = strings.Replace(value, "$", "", 1)
	value = strings.Replace(value, " ", "", 1)
	value = strings.Replace(value, ",", ".", 1)
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

func ConvertStringToFloatScale2(value string) float64 {
	if len(value) == 1 {
		value = fmt.Sprint("0", value)
	}
	value = fmt.Sprint(value[:len(value)-2], ".", value[len(value)-2:])
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

func ConvertStringToTimeLayoutDDMMYYYY(value string) time.Time {
	t, _ := time.Parse(LAYOUT_DDMMYYYY, value)
	return t
}

func ConvertStringToTimeLayoutDD_MM_YYYY(value string) time.Time {
	t, _ := time.Parse(LAYOUT_DD_MM_YYYY, value)
	return t
}

func ConvertStringToTimeLayoutDD_MM_YYYY_HH_MM_SS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_DD_MM_YYYY_HH_MM_SS, value)
	return t
}

func ConvertStringToTimeLayoutYYYYMMDD(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYYMMDD, value)
	return t
}

func ConvertStringToTimeLayoutYYYY_MM_DD(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYY_MM_DD, value)
	return t
}

func ConvertStringToTimeLayoutYYMMDDHHMMSS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYMMDDHHMMSS, value)
	return t
}

func ConvertStringToTimeLayoutYYYYMMDDHHMMSS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYYMMDDHHMMSS, value)
	return t
}

func ConvertStringToTimeLAYOUT_YYYY_MM_DD_HH_MM_SS(date time.Time) string {
	return date.Format(LAYOUT_YYYY_MM_DD_HH_MM_SS)
}

func ConvertStringToTimeLAYOUT_YYYY_MM_DD(date time.Time) string {
	return date.Format(LAYOUT_YYYY_MM_DD)
}

func ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000Z(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYY_MM_DDTHH_MM_SS_000Z, value)
	return t
}

func ConvertStringToTimeLayoutYYYY_MM_DDTHH_MM_SS_000(value string) time.Time {
	t, _ := time.Parse(LAYOUT_YYYY_MM_DDTHH_MM_SS_000, value)
	return t
}

func ConvertStringToTimeLayoutHHMMSS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_HHMMSS, value)
	return t
}

func ConvertStringToTimeLayoutHH_MM_SS(value string) time.Time {
	t, _ := time.Parse(LAYOUT_HH_MM_SS, value)
	return t
}

func ConvertStringToTimeLayoutDDMMYYYYHHMMSS(d time.Time, h time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), h.Hour(), h.Minute(), h.Second(), 0, time.UTC)
}

func ParseBinToHex(s string) string {
	return strconv.FormatInt(int64(ConvertStringToInt(s)), 16)
}

func ConvertStringToInt5Digits(value string) int {
	if len(value) > 5 {
		value = value[:5]
	}
	i, _ := strconv.Atoi(value)
	return i
}

func ConvertFloatToFloatScale2(valor float64) float64 {
	value := strconv.FormatFloat(valor, 'f', 2, 64)
	s, _ := strconv.ParseFloat(value, 64)
	return s
}

func ConvertFloat64ToString(value float64) string {
	s := fmt.Sprintf("%.2f", value)
	return s
}

func ConvertStringToFloat64(value string) float64 {
	s, _ := strconv.ParseFloat(value, 64)
	return s
}
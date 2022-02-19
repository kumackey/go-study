package equinox

import (
	"math"
	"time"
)

const (
	JSTOffset = 9 * 60 * 60
	asiaTokyo = "Asia/Tokyo"
)

var locationJST = time.FixedZone(asiaTokyo, JSTOffset)

// VernalEquinoxDate は春分の日を計算する
func VernalEquinoxDate(year int) time.Time {
	return time.Date(year, time.March, calcVernalEquinoxDate(year), 0, 0, 0, 0, locationJST)
}

// AutumnalEquinoxDate は秋分の日を計算する
func AutumnalEquinoxDate(year int) time.Time {
	return time.Date(year, time.September, calcAutumnalEquinoxDate(year), 0, 0, 0, 0, locationJST)
}

func calcVernalEquinoxDate(year int) int {
	val := calcEquinoxBase(year)

	switch {
	case 1851 <= year && year <= 1899:
		val += 19.8277
	case 1900 <= year && year <= 1979:
		val += 20.8357
	case 1980 <= year && year <= 2099:
		val += 20.8431
	case 2100 <= year && year <= 2150:
		val += 21.8510
	}

	return int(math.Floor(val))
}

func calcAutumnalEquinoxDate(year int) int {
	val := calcEquinoxBase(year)

	switch {
	case 1851 <= year && year <= 1899:
		val += 22.2588
	case 1900 <= year && year <= 1979:
		val += 23.2588
	case 1980 <= year && year <= 2099:
		val += 23.2488
	case 2100 <= year && year <= 2150:
		val += 24.2488
	}

	return int(math.Floor(val))
}

func calcEquinoxBase(year int) float64 {
	return 0.242194*float64(year-1980) - math.Floor(float64(year-1980)/4.0)
}

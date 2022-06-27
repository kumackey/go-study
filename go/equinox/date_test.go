package equinox

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, locationJST)
}

func TestVernalEquinoxDate(t *testing.T) {
	tests := []struct {
		year   int
		output time.Time
	}{
		{2015, d(2015, 3, 21)},
		{2016, d(2016, 3, 20)},
		{2017, d(2017, 3, 20)},
		{2018, d(2018, 3, 21)},
		{2019, d(2019, 3, 21)},
		{2020, d(2020, 3, 20)},
		{2021, d(2021, 3, 20)},
		{2022, d(2022, 3, 21)},
		{2023, d(2023, 3, 21)},
		{2024, d(2024, 3, 20)},
		{2025, d(2025, 3, 20)},
		{2026, d(2026, 3, 20)},
		{2027, d(2027, 3, 21)},
		{2028, d(2028, 3, 20)},
		{2029, d(2029, 3, 20)},
		{2030, d(2030, 3, 20)},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.year), func(t *testing.T) {
			assert.Equal(t, test.output, VernalEquinoxDate(test.year))
		})
	}
}

func TestAutumnalEquinoxDate(t *testing.T) {
	tests := []struct {
		year   int
		output time.Time
	}{
		{2015, d(2015, 9, 23)},
		{2016, d(2016, 9, 22)},
		{2017, d(2017, 9, 23)},
		{2018, d(2018, 9, 23)},
		{2019, d(2019, 9, 23)},
		{2020, d(2020, 9, 22)},
		{2021, d(2021, 9, 23)},
		{2022, d(2022, 9, 23)},
		{2023, d(2023, 9, 23)},
		{2024, d(2024, 9, 22)},
		{2025, d(2025, 9, 23)},
		{2026, d(2026, 9, 23)},
		{2027, d(2027, 9, 23)},
		{2028, d(2028, 9, 22)},
		{2029, d(2029, 9, 23)},
		{2030, d(2030, 9, 23)},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.year), func(t *testing.T) {
			assert.Equal(t, test.output, AutumnalEquinoxDate(test.year))
		})
	}
}

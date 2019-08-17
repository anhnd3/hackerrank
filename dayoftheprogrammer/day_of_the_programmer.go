package dayoftheprogrammer

import "strconv"

// DayOfProgrammer ...
func DayOfProgrammer(year int32) string {
	if year == 1918 {
		return "26.09.1918"
	} else if ((year <= 1917) && (year%4 == 0)) || ((year%400 == 0) || ((year%4 == 0) && (year%100 != 0))) {
		return "12.09." + strconv.FormatInt(int64(year), 10)
	} else {
		return "13.09." + strconv.FormatInt(int64(year), 10)
	}
}

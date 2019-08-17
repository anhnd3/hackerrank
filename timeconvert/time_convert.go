package timeconvert

import "time"

// TimeConversion ...
func TimeConversion(s string) string {
	layout := "03:04:05PM"
	layoutConvert := "15:04:05"
	t, _ := time.Parse(layout, s)
	return t.Format(layoutConvert)
}

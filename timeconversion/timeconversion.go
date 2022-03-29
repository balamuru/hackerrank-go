package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	println(timeConversion("07:01:45PM"))
}

func timeConversion(s string) string {
	// Write your code here
	id := s[8:]
	hms12 := strings.Split(s[:8], ":")
	hr12, _ := strconv.Atoi(hms12[0])
	min, _ := strconv.Atoi(hms12[1])
	sec, _ := strconv.Atoi(hms12[2])

	hms24 := hr12
	if strings.EqualFold(id, "PM") {
		hms24 += 12
		if hms24 == 24 {
			hms24 = 12
		}
	} else if strings.EqualFold(id, "AM") && hms24 == 12 {
		hms24 = 0
	}

	return fmt.Sprintf("%02v:%02v:%02v", hms24, min, sec)

}

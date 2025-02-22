package proxies

import (
	"strconv"
	"strings"
)

var counter = 0

func ResetRenameCounter() {
	counter = 0
}

func GetFlag(countryCode string) string {

	code := strings.ToUpper(countryCode)

	const flagBase = 127397

	first := string(rune(code[0]) + flagBase)
	second := string(rune(code[1]) + flagBase)

	return first + second
}

func Rename(country_code string) string {
	counter++

	if country_code == "" {
		return "UN" + " " + strconv.Itoa(counter)
	}
	country_flag := GetFlag(country_code)

	country_name := country_code

	return country_flag + " " + country_name + " " + strconv.Itoa(counter)
}

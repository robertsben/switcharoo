package main

import (
	"math"
	"strconv"
	"strings"
)

func SanitiseData(charTok string) string {
	charTok = strings.TrimSpace(charTok)
	charTok = strings.Replace(charTok, "\"", "\\\"", -1)
	charTok = strings.Replace(charTok, "\n", "\\n", -1)
	return charTok
}

func IsString(data string) bool {
	_, errb := strconv.ParseBool(data)
	if errb == nil {
		return false
	}
	_, erri := strconv.ParseInt(data, 10, 64)
	if erri == nil {
		return false
	}
	f, errf := strconv.ParseFloat(data, 64)
	if errf == nil {
		if !math.IsNaN(f) {
			return false
		}
	}
	return true
}
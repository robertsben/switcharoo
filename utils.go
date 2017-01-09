package main

import (
	"strings"
)

func SanitiseData(charTok string) string {
	charTok = strings.TrimSpace(charTok)
	charTok = strings.Replace(charTok, "\"", "\\\"", -1)
	charTok = strings.Replace(charTok, "\n", "\\n", -1)
	return charTok
}
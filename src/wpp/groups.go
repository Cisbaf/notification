package wpp

import (
	"strings"
)

const (
	TI = "5521969209913-1548767753"
)

func GetGroup(name string) string {
	_name := strings.ToLower(name)
	if _name == "ti cisbaf" {
		return TI
	}
	return ""
}

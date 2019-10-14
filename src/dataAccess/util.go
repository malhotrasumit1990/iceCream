package dataAccess

import "strings"

var arrayToString = func(input []string) string {
	if len(input) > 0 {
		justString := strings.Join(input, " , ")
		return justString
	}
	return ""
}

var stringToArray = func(input string) []string {
	if len(input) > 0 {
		stringArray := strings.Split(input, ",")
		return stringArray
	}
	return nil
}

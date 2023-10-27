package mygobase

import "net/url"

func validURL(input string) bool {
	_, error := url.Parse(input)
	if error != nil {
		return false
	}
	return true
}

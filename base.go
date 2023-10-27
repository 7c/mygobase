package mygobase

import "net/url"

func ValidURL(input string) bool {
	_, error := url.Parse(input)
	return error == nil
}

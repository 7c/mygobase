package mygobase

import "net/url"

func ValidURL(input string) bool {
	myo, error := url.Parse(input)
	return error == nil && myo.Scheme != "" && myo.Host != ""
}

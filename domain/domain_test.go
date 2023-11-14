package domain

import "testing"

var validDomains []string = []string{
	"golang.dev",
	"golang.net",
	"play.golang.org",
	"gophers.in.space.museum",
	"test.uk",

	"http://golang.dev",
	"http://golang.net",
	"http://play.golang.org",
	"http://gophers.in.space.museum",
	"http://test.uk",

	"https://golang.dev",
	"https://golang.net",
	"https://play.golang.org",
	"https://gophers.in.space.museum",
	"https://test.uk",

	"ftp://golang.dev",
	"ftp://golang.net",
	"ftp://play.golang.org",
	"ftp://gophers.in.space.museum",
	"ftp://test.uk",
}

func TestIsValidDomain(t *testing.T) {
	for _, dom := range validDomains {
		if !IsValidDomain(dom) {
			t.Errorf("Domain '%s' should expected to be VALID", dom)
		}
	}
}

func TestParseDomain(t *testing.T) {
	for _, dom := range validDomains {
		_, err := ParseDomain(dom)
		if err != nil {
			t.Errorf("Domain '%s' should be parseable, got '%s'", dom, err)
		}
		// t.Logf("Domain '%s' parsed as '%+v'", dom, dp)
	}
}

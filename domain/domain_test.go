package domain

import "testing"

var validDomains []string = []string{
	"golang.dev",
	"golang.net",
	"play.golang.org",
	"gophers.in.space.museum",
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
		dp, err := ParseDomain(dom)
		if err != nil {
			t.Errorf("Domain '%s' should be parseable, got '%s'", dom, err)
		}
		if dp.Full != dom {
			t.Errorf("Domain '%s' was parsed but does not match with .Full property", dom)
		}
	}
}

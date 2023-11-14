package domain

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/publicsuffix"
)

type Managed string
type DomainParsed struct {
	Tld     string
	Domain  string
	Full    string
	Sub     string
	Manager Managed
}

const (
	ICAAN     Managed = "ICAAN"
	Unmanaged Managed = "UNMANAGED"
	Private   Managed = "PRIVATE"
)

var (
	ERRInvalidDomain  = errors.New("INVALID DOMAIN")
	ERRNotIcaanDomain = errors.New("NOT AN ICAAN DOMAIN")
)

func pslManager(domain string) (Managed, string) {
	eTLD, icann := publicsuffix.PublicSuffix(domain)
	manager := Unmanaged
	if icann {
		manager = ICAAN
	} else if strings.IndexByte(eTLD, '.') >= 0 {
		manager = Private
	}
	return manager, eTLD
}

var domainRegexp = regexp.MustCompile(`^(?i)[a-z0-9-]+(\.[a-z0-9-]+)+\.?$`)

// IsValidDomain returns true if the domain is valid.
//
// It uses a simple regular expression to check the domain validity.
func IsValidDomain(domain string) bool {
	return domainRegexp.MatchString(domain)
}

// parsed ICAAN domains
func ParseDomain(domain string) (*DomainParsed, error) {
	domain = strings.TrimSpace(domain)

	// lets support urls as well
	u, err1 := url.Parse(domain)
	if err1 != nil {
		return nil, ERRInvalidDomain
	}
	if u.Scheme != "" && u.Host != "" {
		domain = u.Host
	}

	if man, eTLD := pslManager(domain); man == ICAAN {
		dp := new(DomainParsed)
		etld_len := len(strings.Split(eTLD, "."))
		dp.Tld = eTLD
		dp.Manager = ICAAN

		parts := strings.Split(domain, ".")
		parts = parts[:len(parts)-etld_len] // remove the eTLD part like com.au
		if len(parts) == 0 {
			return nil, ERRInvalidDomain
		}
		// litter.Dump(parts)
		dp.Domain = fmt.Sprintf(`%s.%s`, strings.Join(parts[len(parts)-1:], "."), eTLD)
		if !IsValidDomain(dp.Domain) {
			return nil, ERRInvalidDomain
		}

		dp.Sub = strings.Join(parts[:len(parts)-1], ".")

		// dp.Domain = fmt.Sprintf(`%s`, strings.Join(parts[len(parts)-1:], "."))

		dp.Full = fmt.Sprintf(`%s.%s`, strings.Join(parts, "."), eTLD)
		// litter.Dump(dp)
		return dp, nil
	}
	return nil, ERRNotIcaanDomain
}

package email

import (
	"regexp"
	"strings"
)

// Regular expression from WebCore's HTML5 email input: http://goo.gl/7SZbzj
var emailRegexp = regexp.MustCompile("(?i)" + // case insensitive
	"^[a-z0-9!#$%&'*+/=?^_`{|}~.-]+" + // local part
	"@" +
	"[a-z0-9-]+(\\.[a-z0-9-]+)*$") // domain part

// IsValidEmail returns true if the given string is a valid email address.
//
// It uses a simple regular expression to check the address validity.
func IsValidEmail(email string) bool {
	if len(email) > 254 {
		return false
	}
	return emailRegexp.MatchString(email)
}

// splitEmail splits email address into local and domain parts.
// The last returned value is false if splitting fails.
func splitEmail(email string) (local string, domain string, ok bool) {
	parts := strings.Split(email, "@")
	if len(parts) < 2 {
		return
	}
	local = parts[0]
	domain = parts[1]
	// Check that the parts contain enough characters.
	if len(local) < 1 {
		return
	}
	if len(domain) < len("x.xx") {
		return
	}
	return local, domain, true
}

// NormalizeEmail returns a normalized email address.
// It returns an empty string if the email is not valid.
func NormalizeEmail(email string) string {
	// Trim whitespace.
	email = strings.TrimSpace(email)
	// Make sure it is valid.
	if !IsValidEmail(email) {
		return ""
	}
	// Split email into parts.
	local, domain, ok := splitEmail(email)
	if !ok {
		return ""
	}
	// Remove trailing dot from domain.
	domain = strings.TrimRight(domain, ".")
	// Convert domain to lower case.
	domain = strings.ToLower(domain)
	// Combine and return the result.
	return local + "@" + domain
}

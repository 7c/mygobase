# mygobase

## import
`go get -u github.com/7c/mygobase`


## base functions
```
ValidURL(url string) bool
ValidIP(ip string) bool 
ValidIP4(ip string) bool 
ValidIP6(ip string) bool 
FileExists(filename string) bool
LockPort(port int) bool 
```

## /domain
### ParseDomain(domain string) (*DomainParsed, error)

## /email
### NormalizeEmail(email string) string
### IsValidEmail(email string) bool
## /vault
### Model: ModelVault,.init,.KVv1,.AsSqlUri

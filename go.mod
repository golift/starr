module golift.io/starr

go 1.17

require (
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // publicsuffix, cookiejar.
	golift.io/datacounter v1.0.3 // Counts bytes read from starr apps.
)

// All of this is for the tests.
require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/golang/mock v1.6.0
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.7.0
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
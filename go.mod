module golift.io/starr

go 1.17

require (
	golang.org/x/net v0.0.0-20220114011407-0dd24b26b47d // publicsuffix, cookiejar.
)

// All of this is for the tests.
require (
	github.com/golang/mock v1.6.0
	github.com/stretchr/testify v1.7.0
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

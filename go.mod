module golift.io/starr

go 1.25.7

toolchain go1.26.3

require golang.org/x/net v0.49.0 // publicsuffix, cookiejar.

// All of this is for the tests.
require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.11.1 // assert!
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

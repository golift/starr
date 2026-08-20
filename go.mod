module golift.io/starr

go 1.25.7

toolchain go1.27.0

require golang.org/x/net v0.58.0 // publicsuffix, cookiejar.

// All of this is for the tests.
require (
	github.com/stretchr/testify v1.12.1 // assert!
	go.yaml.in/yaml/v3 v3.0.5 // indirect
)

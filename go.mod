module golift.io/starr

go 1.17

require (
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // publicsuffix, cookiejar.
	golift.io/datacounter v1.0.3 // Counts bytes read from starr apps.
)

// All of this is for the tests.
require (
	github.com/stretchr/testify v1.7.0 // assert!
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

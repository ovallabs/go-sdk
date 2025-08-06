module github.com/ovalfi/go-sdk

go 1.17

//replace github.com/go-resty/resty/v2 => gopkg.in/resty.v1 v1.12.0

require (
	github.com/go-resty/resty/v2 v2.11.0
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/zerolog v1.28.0
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

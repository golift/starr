all:
	@echo "try: make test"

test: lint nopollution
	go test -race -covermode=atomic ./...
	# Test 32 bit OSes.
	GOOS=linux GOARCH=386 go build .
	GOOS=freebsd GOARCH=386 go build .

lint:
	# Test lint on four platforms.
	GOOS=linux golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct
	GOOS=darwin golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct
	GOOS=windows golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct
	GOOS=freebsd golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct

nopollution:
	# Avoid cross pollution.
	grep -riE 'readar|radar|sonar|prowl' lidarr   || exit 0 && exit 1
	grep -riE 'readar|sonar|lidar|prowl' radarr   || exit 0 && exit 1
	grep -riE 'radar|sonar|lidar|prowl'  readarr  || exit 0 && exit 1
	grep -riE 'readar|radar|lidar|prowl' sonarr   || exit 0 && exit 1
	grep -riE 'readar|radar|lidar|sonar' prowlarr || exit 0 && exit 1

generate:
	go generate ./...

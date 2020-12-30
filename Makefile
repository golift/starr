all:
	@echo "try make test"

test: lint nopollution
	go test -race -covermode=atomic ./...
	# Test 32 bit OSes.
	GOOS=linux GOARCH=386 go build .
	GOOS=freebsd GOARCH=386 go build .

lint:
	GOOS=linux golangci-lint run --enable-all -D makezero,forbidigo
	GOOS=darwin golangci-lint run --enable-all -D makezero,forbidigo
	GOOS=windows golangci-lint run --enable-all -D makezero,forbidigo
	GOOS=freebsd golangci-lint run --enable-all -D makezero,forbidigo

nopollution:
	# Avoid cross pollution.
	grep -riE 'readar|sonar|lidar' radarr  || exit 0 && exit 1
	grep -riE 'radar|sonar|lidar'  readarr || exit 0 && exit 1
	grep -riE 'readar|radar|lidar' sonarr  || exit 0 && exit 1
	grep -riE 'readar|radar|sonar' lidarr  || exit 0 && exit 1

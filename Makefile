all:
	@echo "try: make test"

test: lint nopollution
	go test -race -covermode=atomic ./...
	# Test 32 bit OSes.
	GOOS=linux GOARCH=386 go build .
	GOOS=freebsd GOARCH=386 go build .

lint:
	# Test lint on four platforms.
	GOOS=linux golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct,dupl
	GOOS=darwin golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct,dupl
	GOOS=windows golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct,dupl
	GOOS=freebsd golangci-lint run --enable-all -D maligned,scopelint,interfacer,golint,tagliatelle,exhaustivestruct,dupl

# Some of these are borderline. For instance "edition" shows up in radarr payloads. "series" shows up in Readarr, "author" in Sonarr, etc.
# If these catch legitimate uses, just remove the piece that caught it.
nopollution:
	# Avoid cross pollution.
	grep -riE 'readar|radar|sonar|prowl|series|episode|author|book|edition|movie|v3' lidarr   || exit 0 && exit 1
	grep -riE 'readar|sonar|lidar|prowl|series|episode|author|book|artist|album|v1' radarr   || exit 0 && exit 1
	grep -riE 'radar|sonar|lidar|prowl|episode|movie|artist|album|v3'  readarr  || exit 0 && exit 1
	grep -riE 'readar|radar|lidar|prowl|book|edition|movie|artist|album|v1' sonarr   || exit 0 && exit 1
	grep -riE 'readar|radar|lidar|sonar|series|episode|author|book|edition|movie|artist|album|track|v3' prowlarr || exit 0 && exit 1

generate:
	go generate ./...

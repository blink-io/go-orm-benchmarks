
.PHONY: tidy
# Run go mod tidy
tidy:
	go mod tidy -v -e

.PHONY: upgrade
# Upgrade packages
upgrade:
	go get -u -v -t ./...
	$(MAKE) tidy

.PHONY: build
# Upgrade packages
build:
	go build -v ./...
	$(MAKE) tidy


.PHONY: gen-bob
gen-bob:
	go run github.com/stephenafamo/bob/gen/bobgen-psql@latest -c ./bench/bob/bobgen.yaml
BASE_PKG=github.com/akm/reflectjson
VERSION ?= $(shell cat ./VERSION)

.PHONY: test
test:
	go test . && \
	go test $(BASE_PKG)/typedict

.PHONY: build
build:
	go build .

.PHONY: version
version:
	@echo $(VERSION)	

.PHONY: git_guard
git_guard:
	@git diff --exit-code

.PHONY: release
release: git_guard
	git tag $(VERSION) && \
	git push origin $(VERSION)

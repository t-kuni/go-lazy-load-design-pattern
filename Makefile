.PHONY: generate
generate:
	go generate -x ./...

.PHONY: test
test: generate
	gotestsum --hide-summary=skipped -- ./... -v

.PHONY: install-tools
install-tools:

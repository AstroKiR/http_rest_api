.PHONY: build
build:
	go build -v cmd/apiserver/main.go

.PHONY: run
run:
	go run -v cmd/apiserver/main.go

.DEFAULT_GOAL: build
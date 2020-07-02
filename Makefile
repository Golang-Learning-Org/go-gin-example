run:
	go run main.go

build: docs
	go build main.go

docs:
	swag init

.PHONY: run build docs
run:
	go run main.go

build: docs
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-gin-example .

docs:
	swag init

.PHONY: run build docs

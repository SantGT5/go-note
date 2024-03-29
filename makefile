#  Variables
APP_NAME=gopportunities

# Tasks
.PHONY: default
default: run

.PHONY: docs
run: docs
	@go run main.go

.PHONY: build
build:
	@go build -o ${APP_NAME} main.go

.PHONY: test
test:
	go test

.PHONY: docs
docs:
	@swag init

.PHONY: clean
clean:
	@rm -f ${APP_NAME}
	@rm -rf ./docs

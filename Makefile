BUILD_DIR := build

deps::
	go mod download

build:: deps
	go build -o $(BUILD_DIR)/unsubme ./cmd/unsubme/main.go

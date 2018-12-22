BUILD_DIR := build

build::
	go build -o $(BUILD_DIR)/unsubme ./cmd/unsubme/main.go

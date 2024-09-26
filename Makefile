APP_NAME = server

.PHONY: run

run:
	go run ./cmd/$(APP_NAME)/

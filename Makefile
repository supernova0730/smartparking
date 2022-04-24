export BIN_DIR=./cmd/bin
export BIN_NAME=smartparking

build:
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(BIN_NAME)

run: build
	$(BIN_DIR)/$(BIN_NAME) start

migrate: build
	$(BIN_DIR)/$(BIN_NAME) migrate

deploy: build
	ssh $(USER)@$(HOST) "echo $(PASS) sudo -S systemctl stop $(BIN_NAME) && rm -rf /home/$(USER)/$(BIN_NAME)"
	scp $(BIN_DIR)/$(BIN_NAME) $(USER)@$(HOST):/home/$(USER)
	ssh $(USER)@$(HOST) "echo $(PASS) sudo -S systemctl start $(BIN_NAME)"

PROGRAM := navi
VERSION := 0.0.0

BASE := $(shell pwd)

CMD_PATH := $(BASE)/app/run

BIN_DIR := $(BASE)/bin
BIN_FILE := $(BIN_DIR)/$(PROGRAM)


BUILD_PKG := $(shell head -1 $(BASE)/go.mod | cut -d ' ' -f 2)

BUILD_DATE := $(shell date -u +%Y-%m-%d.%H:%M:%S-%Z)
GIT_COMMIT := $(shell git rev-parse HEAD)

LDFLAGS :=  -ldflags "\
	-X $(BUILD_PKG).version=$(VERSION) \
	-X $(BUILD_PKG).date=$(BUILD_DATE) \
	-X $(BUILD_PKG).commit=$(GIT_COMMIT)" \

$(PROGRAM):
	go build $(LDFLAGS) -o $(BIN_FILE) $(CMD_PATH)

clean-$(PROGRAM):
	rm $(BIN_AGENT)

clean: clean-$(PROGRAM)
	rmdir $(BIN_DIR)

all: $(PROGRAM)

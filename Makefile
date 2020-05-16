PROGRAM := navi
VERSION := 0.0.0

BASE := $(shell pwd)

BIN_DIR := $(BASE)/bin

BIN_FILE := $(BIN_DIR)/$(PROGRAM)

CMD_PATH := $(BASE)/app/run


BUILD_DATE := $(shell date -u +%Y-%m-%d.%H:%M:%S-%Z)
GIT_COMMIT := $(shell git rev-parse HEAD)


VERSION_PKG := $(shell head -1 go.mod | cut -d ' ' -f 2)

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

EXTENSION := "issue-stats"

MAKEFLAGS += --no-print-directory

# Version information
VERSION ?= $(shell git describe --tags --always --dirty)

# LDFLAGS
LDFLAGS := -X github.com/shufo/gh-issue-stats/cmd.Version=$(VERSION)

build:
	@go build -ldflags "$(LDFLAGS)"
run:
	@make build
	@gh $(EXTENSION)
install:
	@gh extension install . --force

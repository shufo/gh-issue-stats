EXTENSION := "issue-stats"

MAKEFLAGS += --no-print-directory

build:
	@go build
run:
	@make build
	@gh $(EXTENSION)
install:
	@gh extension install . --force

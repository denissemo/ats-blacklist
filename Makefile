.PHONY: build
.PHONY: run

## build: Build and run server
build:
	go build -o ./build -v .
	./build/ats-blacklist # run App

run:
	go build -o ./build -v .
	pm2 start ./build/ats-blacklist

help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

.DEFAULT_GOAL := build
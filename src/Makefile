# global command
.PHONY: all
#all: clean dependencies platform-itsm-sync-service  pack-service
all: clean dependencies build compile

.PHONY: clean
clean:
	@echo "\n*********Cleaning in progress*************"
	rm -rf cover-all.out
	rm -rf cover-all.html

.PHONY: dependencies
dependencies:
	@echo "\n*********Installing dependencies*************"
	rm -f glide.lock
	glide cc
	glide install


.PHONY: build
build:
	go build -o bin/main main.go

.PHONY: compile
compile:
	@echo "\n*************Generating Binary*************"
	GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/main-linux-amd64 main.go	


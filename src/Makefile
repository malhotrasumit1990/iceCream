# global command
.PHONY: all
#all: clean dependencies platform-itsm-sync-service  pack-service
all: clean dependencies build compile test

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

#Target to run test suits
.PHONY: test
test:
	@echo "\n*************Executing testcases*************"
	go test -coverprofile ./web/controller/cover.out -covermode=count ./web/controller
	go test -coverprofile ./dataAccess/cover.out -covermode=count ./dataAccess

.PHONY: compile
compile:
	@echo "\n*************Generating Binary*************"
	GOOS=linux GOARCH=amd64 go build -o $(GOPATH)/bin/main-linux-amd64 main.go	


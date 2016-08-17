CC=go
export GOPATH=${PWD}

debug: 
	$(CC) run src/server.go config.yml

build: clean
	$(CC) build src/server.go
	mv server bin/

deps:
	$(CC) get gopkg.in/yaml.v2
	$(CC) get github.com/stianeikeland/go-rpio

.PHONY: clean
clean:
	rm -rf bin/*

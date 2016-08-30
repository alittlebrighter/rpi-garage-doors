CC=go

debug: 
	$(CC) run src/server.go config.yaml

build: clean
	$(CC) build src/server.go
	mv server bin/

deps:
	$(CC) get gopkg.in/yaml.v2
	$(CC) get github.com/stianeikeland/go-rpio

.PHONY: clean
clean:
	rm -rf bin/*

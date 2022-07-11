
build:
	sudo docker build -t pipelines:latest .

test:
	go test ./test/...

.PHONY: build test
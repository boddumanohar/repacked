build:
	go build -o main

docker:
	docker build -t pack:latest .

run: build
	./main

dockerrun: docker
	docker run --rm -p 8080:8080 pack:latest

test:
	go test -v ./...

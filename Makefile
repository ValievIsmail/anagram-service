go: 
	go run cmd/*.go

test:
	go test -v ./...
	
build:
	go build -o ./anagram-service cmd/*.go

docker-build:
	docker build -t anagram-service .

docker-run:
	docker run -d -p 8080:8080 anagram-service
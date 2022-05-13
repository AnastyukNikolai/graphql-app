build:
	gofmt -s -w ./
	goimports -w -d ./
# 	golangci-lint run ./...
	go generate ./...
	docker-compose build graphql-app

run:
	docker-compose up graphql-app

test:
	go test -v ./...

download:
	go mod download
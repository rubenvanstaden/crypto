fmt:
	go mod tidy -compat=1.17
	gofmt -l -s -w .

build:
	go build -o ./bin/crypto ./cmd/*.go

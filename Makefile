fmt:
	go mod tidy -compat=1.17
	gofmt -l -s -w .

build:
	go build -o ./bin/nkey ./cmd/*.go

install:
	cp -f ./bin/nkey $(HOME)/go/bin/

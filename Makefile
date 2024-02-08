TEMP_DIR=./tmp/
BINARY_DIR=./bin/
BINARY_NAME=cmus-status

all: test clean init build

clean:
	go clean
	rm -rf ${TEMP_DIR}
	rm -rf ${BINARY_DIR}

init:
	go mod tidy
	go mod download

test:
	go test ./...

test_coverage:
	mkdir -p ${TEMP_DIR}
	go test -coverprofile ${TEMP_DIR}cover.out ./...
	go tool cover -html=${TEMP_DIR}cover.out -o ${TEMP_DIR}cover.html
	firefox ${TEMP_DIR}cover.html

build:
	mkdir -p ${BINARY_DIR}
	go build -o ${BINARY_DIR}${BINARY_NAME} ./cmd/cmus-status/main.go

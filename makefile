BINARY_NAME=main

run: 
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME} main.go

clean:
	go clean
BINARY_NAME=goco
OUT=./out/${BINARY_NAME}

build:
	mkdir -p out
	GOARCH=amd64 GOOS=darwin go build -o ${OUT}-darwin ./cmd/root.go
	GOARCH=amd64 GOOS=linux go build -o ${OUT}-linux ./cmd/root.go
	GOARCH=amd64 GOOS=windows go build -o ${OUT}-windows ./cmd/root.go

run: build
	${OUT}-darwin  # or ${OUT}-linux depending on your OS

clean:
	go clean
	rm -f ${OUT}-darwin ${OUT}-linux ${OUT}-windows

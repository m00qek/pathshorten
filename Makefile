all: run

clean:
	@rm bin/ -rf

run:
	@go run *.go

build: clean
	@go build -o bin/pathshorten

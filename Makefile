all: run

clean:
	@rm bin/ -rf

run:
	@go run *[^_test].go ${PWD}

build: clean
	@go build -o bin/pathshorten

test:
	@go test

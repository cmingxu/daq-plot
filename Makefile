default: all

all: build
	bin/dplot

build:
	@mkdir -p bin
	@go build -o bin/dplot ./...
	
clean:
	@rm -rf bin


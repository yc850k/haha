all: build

build: build-server

build-server:
	go build -o bin/configServer ./server.go

clean:
	@rm -rf bin
	@if [ -d test ]; then cd test && rm -f *.out *.log *.rdb; fi

gotest:
	go test ./pkg/... ./... -race

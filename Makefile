build:
	go build -v

fmt:
	gofmt -w -s *.go

test:
	go test -v -cover

bench:
	go test -bench=. -test.benchmem

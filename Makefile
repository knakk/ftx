all: build

build:
	export GOBIN=$(shell pwd)
	go build

profile:
	go test -run none -bench . -benchtime 4s -cpuprofile=prof.out
	go tool pprof ./ftx.test ./prof.out

clean:
	go clean
	rm -f ftx.test
	rm -f prof.out
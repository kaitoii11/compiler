
icc:
	go build -o icc cmd/icc/main.go

test: icc
	./test.sh

clean:
	go clean
	rm -r icc

.PHONY: test clean
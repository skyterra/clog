.PHONY: test cover clean

test:
	go test ./...

cover:
	go test ./... -coverprofile coverage.out

clean:
	rm -f coverage.out
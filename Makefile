.PHONY: clean lint security critic test

clean:
	rm -rf ./tmp coverage.out

lint:
	golangci-lint run ./...

security:
	gosec -quiet ./...

critic:
	gocritic check -enableAll ./...

test: clean lint security critic
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

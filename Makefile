.PHONY: lint test run

run: 
	go install ./src/ && go run ./src/

test: 
	go get -u golang.org/x/lint/golint
	CGO_ENABLED=1 go test -v -coverprofile=coverage.txt -covermode=atomic -count=1 ./src/...
	cd ./src && golint
	
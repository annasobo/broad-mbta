.PHONY: lint test run

run: 
	go install ./src/ && go run ./src/

test: 
	go install ./src/ && go test ./src/
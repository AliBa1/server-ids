.PHONY: coverage-cli coverage-html

coverage-cli:
	go test ./... -coverprofile=cover.out
	go tool cover -func=cover.out

coverage-html:
	go test ./... -coverprofile=cover.out
	go tool cover -html=cover.out
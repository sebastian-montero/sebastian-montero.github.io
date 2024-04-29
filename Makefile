.PHONY: build clean

build:
	go build -o bin/ssg cmd/main.go

# # Run test
# test:
# 	go test ./...

clean:
	rm -rf bin/*

run:
	go run cmd/main.go -md ./content/page.md -tmpl ./templates/template.html -out ./public/index.html -title "My Page"

fmt:
	goimports -w ./cmd
	gofmt -w ./cmd

lint:
	staticcheck ./cmd

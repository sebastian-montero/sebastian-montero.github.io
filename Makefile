.PHONY: build clean

build:
	go build -o bin/ssg cmd/ssg/main.go

# # Run test
# test:
# 	go test ./...

clean:
	rm -rf bin/*

run:
	go run cmd/ssg/main.go -md ./content/page.md -tmpl ./templates/template.html -out ./public/index.html -title "My Page"

fmt:
	gofmt -w ./cmd

lint:
	staticcheck ./cmd

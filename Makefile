.PHONY: build clean

all: build exec

build:
	@echo "Building..."
	mkdir -p bin
	go build -o bin/ssg cmd/main.go

clean:
	@echo "Running cleanup..."
	rm -rf bin/*

run:
	@echo "Run..."
	go run cmd/main.go -md ./content/about.md -tmpl ./templates/about_template.html -out ./public/index.html -title "SM"

exec:
	@echo "Run..."
	./bin/ssg -md ./content/about.md -tmpl ./templates/about_template.html -out ./public/index.html -title "SM"

fmt:
	@echo "Running gofmt..."
	goimports -w ./cmd
	gofmt -w ./cmd
	goimports -w ./pkg
	gofmt -w ./pkg

lint:
	@echo "Running lint..."
	staticcheck ./cmd

# # Run test
# test:
# 	go test ./...
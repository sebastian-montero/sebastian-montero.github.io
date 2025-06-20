.PHONY: build clean fmt exec all

all: fmt build exec

build:
	@echo "Building..."
	mkdir -p bin
	go build -o bin/ssg cmd/main.go

clean:
	@echo "Running cleanup..."
	rm -rf bin/*

run:
	@echo "Run..."
	go run cmd/main.go -md ./content/about.md -tmpl ./templates/about_template.html -out ./index.html -title "SM"

run-gallery:
	@echo "Run..."
	go run cmd/gallery/gallery.go -img content/brutalism/data.yaml -tmpl templates/brutalist_template.html -out content/brutalism/index.html

exec:
	@echo "Executing..."
	./bin/ssg -md ./content/about.md -tmpl ./templates/about_template.html -out ./index.html -title "SM"

fmt:
	@echo "Running gofmt..."
	goimports -w ./cmd
	gofmt -w ./cmd
	goimports -w ./pkg
	gofmt -w ./pkg

lint:
	@echo "Running lint..."
	staticcheck ./cmd

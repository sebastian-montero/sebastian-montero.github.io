.PHONY: build clean

build:
	go build -o bin/ssg cmd/main.go

# # Run test
# test:
# 	go test ./...

clean:
	rm -rf bin/*

run:
	go run cmd/main.go -md ./content/about.md -tmpl ./templates/about_template.html -out ./public/index.html -title "SM"

exec:
	./bin/ssg -md ./content/about.md -tmpl ./templates/about_template.html -out ./public/index.html -title "SM"

fmt:
	goimports -w ./cmd
	gofmt -w ./cmd
	goimports -w ./pkg
	gofmt -w ./pkg

lint:
	staticcheck ./cmd

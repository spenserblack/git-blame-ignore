SOURCES := $(wildcard *.go) $(wildcard pkg/*/*.go)

dist/git-blame-ignore: $(SOURCES)
	go build -trimpath -o ./dist/ ./...

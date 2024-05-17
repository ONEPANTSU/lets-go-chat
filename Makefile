PHONY:
.SILENT:

build-app:
	go build -ldflags="-s -w" -o ./.bin/main cmd/app/main.go

run: build-app
	./.bin/main
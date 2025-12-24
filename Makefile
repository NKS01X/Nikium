file ?= example.nik

run:
	go run main.go $(file)

build:
	go build -o nikium.exe

.PHONY: run build

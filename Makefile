FILE ?= example.nik

run: build
	./nikium $(FILE)

build:
	go build -o nikium main.go

.PHONY: run build

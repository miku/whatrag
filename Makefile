SHELL := /bin/bash

llman: llman.go
	go build -o llman llman.go

.PHONY: clean
clean:
	rm -f llman

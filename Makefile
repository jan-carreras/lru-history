#!/usr/bin/env make

test:
	go test ./...

install:
	go install h.go

fmt:
	go fmt ./...

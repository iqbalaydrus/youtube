#!/usr/bin/env bash
rm -f output.csv plot-golang.png
go mod download
go build -o dist/goroutine 20240613-goroutine
psrecord './dist/goroutine' --include-children --plot plot-golang.png

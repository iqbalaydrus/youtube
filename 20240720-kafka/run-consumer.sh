#!/usr/bin/env bash
mkdir -p dist
go mod download
go build -o dist/kafka 20240720-kafka
./dist/kafka consumer test1

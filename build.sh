#!/bin/bash

go build -o bowerbird cmd/main.go

sudo cp bowerbird /usr/local/bin/

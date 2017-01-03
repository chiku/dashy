# Makefile
#
# Author::    Chirantan Mitra
# Copyright:: Copyright (c) 2015-2017. All rights reserved
# License::   MIT

MAKEFLAGS += --warn-undefined-variables
SHELL := bash

.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
.DELETE_ON_ERROR:
.SUFFIXES:

ifndef GOPATH
$(error GOPATH not set)
endif

MKDIR = mkdir -p
RM = rm -rvf
GO = go

sources := $(wildcard app/*.go)
app = ./app
main = main.go
out = out
packages = . $(app)
binary = $(out)/dashy
coverage = $(out)/coverage
coverage_out = $(coverage)/coverage.out
coverage_html = $(coverage)/coverage.html

all: fmt vet test compile
.PHONY: all

fmt:
	${GO} fmt $(packages)
.PHONY: fmt

vet:
	${GO} vet $(packages)
.PHONY: vet

test: $(coverage_html)
.PHONY: test

compile: $(binary)
.PHONY: compile

$(binary):
	${GO} build -o $(binary) $(main)

clean:
	${RM} $(binary) $(coverage)
.PHONY: clean

$(coverage_out): $(sources)
	${MKDIR} $(coverage)
	${GO} test $(app) -coverprofile=$(coverage_out)

$(coverage_html): $(coverage_out)
	${GO} tool cover -func=$(coverage_out)
	${GO} tool cover -html=$(coverage_out) -o $(coverage_html)

#!/bin/bash

# Goのテストを行うやつ

abspath="$(realpath .)"
go test -v ./src/... -test.config "$abspath/db/sqlboiler.toml" $@

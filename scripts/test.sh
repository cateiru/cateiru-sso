#!/bin/bash

# Goのテストを行うやつ

export RECAPTCHA_SECRET=secret
export MAILGUN_SECRET=secret
export FASTLY_API_TOKEN=token
export STORAGE_EMULATOR_HOST=localhost:4443

abspath="$(realpath .)"
go test ./src/... -test.config "$abspath/db/sqlboiler.toml" $@

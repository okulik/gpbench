package main

import (
	"testing"
)

func TestAssert_WithTooManyWorkers(t *testing.T) {
	clip := cliParamsType{workers: 300}
	compare(clip.assert().Error(), "number of workers is limited to 200", t)
}

func TestInitCliParams_InvalidInputFile(t *testing.T) {
	clip := cliParamsType{params: "no_such_file"}
	compare(clip.assert().Error(), "stat no_such_file: no such file or directory", t)
}

func TestConnectionString(t *testing.T) {
	clip := cliParamsType{
		host:     "pgdb001.pgdb.com",
		port:     1234,
		database: "postgres",
		user:     "postgres",
		password: "password"}
	compare(clip.connectionString(), "user=postgres dbname=postgres host=pgdb001.pgdb.com port=1234 password=password sslmode=disable", t)
}

func TestConnectionString_NoPassword(t *testing.T) {
	clip := cliParamsType{
		host:     "pgdb001.pgdb.com",
		port:     1234,
		database: "postgres",
		user:     "postgres"}
	initCliParams()
	compare(clip.connectionString(), "user=postgres dbname=postgres host=pgdb001.pgdb.com port=1234 sslmode=disable", t)
}

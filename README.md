# go-oauth2-mysql
MySQL storage for OAuth 2.0  Provides both client and token store.

[![Build][Build-Status-Image]][Build-Status-Url] [![Codecov][codecov-image]][codecov-url] [![ReportCard][reportcard-image]][reportcard-url] [![GoDoc][godoc-image]][godoc-url] [![License][license-image]][license-url]

## Install

``` bash
$ go get -u -v github.com/vnzernov/go-oauth2-mysql
```

## Usage

``` go
package main

import (
	"github.com/vnzernov/go-oauth2-mysql"
	"gopkg.in/oauth2.v3/manage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	manager := manage.NewDefaultManager()
	dsn := "root:123456@tcp(127.0.0.1:3306)/myapp_test?charset=utf8"
	// use mysql token store
	store := mysql.NewDefaultStore(
		mysql.NewConfig(dsn),
	)

	defer store.Close()
	// use mysql client store
	clientStore := mysql.NewClientDefaultStore(
		mysql.NewConfig(dsn),
	)

	defer clientStore.Close()

	clientStore.Set(&models.Client{
		ID:     "000000",
		Secret: "999999",
		Domain: "http://localhost",
	})

	manager.MapTokenStorage(store)
	manager.MapClientStorage(clientStore)

// ...
}

```
Credits

Based on https://github.com/go-oauth2/mysql/

## MIT License

[Build-Status-Url]: https://travis-ci.org/vnzernov/go-oauth2-mysql
[Build-Status-Image]: https://travis-ci.org/vnzernov/go-oauth2-mysql.svg?branch=master
[codecov-url]: https://codecov.io/gh/vnzernov/go-oauth2-mysql
[codecov-image]: https://codecov.io/gh/vnzernov/go-oauth2-mysql/branch/master/graph/badge.svg
[reportcard-url]: https://goreportcard.com/report/github.com/vnzernov/go-oauth2-mysql
[reportcard-image]: https://goreportcard.com/badge/github.com/vnzernov/go-oauth2-mysql
[godoc-url]: https://godoc.org/github.com/vnzernov/go-oauth2-mysql
[godoc-image]: https://godoc.org/github.com/vnzernov/go-oauth2-mysql?status.svg
[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg

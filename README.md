# go-oauth2-mysql
MySQL storage for OAuth 2.0  Provides both client and token store.

[![License][license-image]][license-url]

## Install

``` bash
$ go get -u -v gopkg.in/go-oauth2/mysql.v3
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

```
Copyright (c) 2020
```

[license-url]: http://opensource.org/licenses/MIT
[license-image]: https://img.shields.io/npm/l/express.svg

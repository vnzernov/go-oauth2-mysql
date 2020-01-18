package mysql

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/oauth2.v3/models"
)

func TestClientStore(t *testing.T) {
	Convey("Test mysql client store", t, func() {
		Convey("Test error connect", func() {
			So(func() { NewClientDefaultStore(NewConfig("")) }, ShouldPanic)
		})
		envDsn, ok := os.LookupEnv("MYSQL_DSN")
		if ok {
			dsn = envDsn
		}
		clientStore := NewClientDefaultStore(NewConfig(dsn))
		Convey("Test create client", func() {
			err := clientStore.Create(&models.Client{
				ID:     "000000",
				Secret: "999999",
				Domain: "http://localhost",
			})
			So(err, ShouldBeNil)
		})
		Convey("Test get client", func() {
			client, err := clientStore.GetByID("000000")
			So(err, ShouldBeNil)
			So(client.GetID(), ShouldEqual, "000000")
		})
		Convey("Test get client (client not found", func() {
			client, err := clientStore.GetByID("000001")
			So(err, ShouldBeNil)
			So(client, ShouldBeNil)
		})
		Convey("Test delete client", func() {
			err := clientStore.Delete("000000")
			So(err, ShouldBeNil)
		})
		Convey("Test error delete client", func() {
			err := clientStore.Delete("000001")
			So(err, ShouldBeNil)
		})
		Convey("Test close client", func() {
			err := clientStore.Close()
			So(err, ShouldBeNil)
		})
	})
}

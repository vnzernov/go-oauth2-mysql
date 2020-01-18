package mysql

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/oauth2.v3/models"
)

const (
//dsn = "root:secret@tcp(172.17.0.3:3306)/yamalhe?charset=utf8"
)

func TestClientStore(t *testing.T) {
	Convey("Test mysql client store", t, func() {
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
	})
}


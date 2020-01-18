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
		Convey("Test delete client", func() {
			err := clientStore.Delete("000000")
			So(err, ShouldBeNil)
		})
	})
}

/*
import (
	"gopkg.in/oauth2.v3/models"
	"testing"
)

const (
	dbTypeClient = "sqlite3"
	//dsn = "file::memory:?cache=shared"
	dsnClient = "/home/zvn/test.db"
)

// TestClientStore -
func TestClientStore(t *testing.T) {
	info := &models.Client{
		ID:     "1",
		Secret: "111111",
		Domain: "http://localhost",
		UserID: "1",
	}
	store := NewClientStore(NewConfig(dsnClient, dbTypeClient, ""), 1)
	err := store.Create(info)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	client, err := store.GetByID("1")
	if err != nil {
		t.Error(err)
	}
	if client != info {
		t.Error("Error read client info")
	}
}


	Convey("Test client store", t, func() {
		store := NewClientStore(NewConfig(dsn, dbType, ""), 1)

		Convey("Test client store", func() {
			info := &models.Client{
				ID:     "1",
				Secret: "111111",
				Domain: "http://localhost",
				UserID: "1",
			}
			err := store.Create(info)
			if err != nil {
				t.Error(err)
			}
			//So(err, ShouldBeNil)
		})
		Convey("Test get client store", func() {
			client, err := store.GetByID("1")
			if err != nil {
				t.Error(err)
			}
			if client !=
			//So(err, ShouldBeNil)

		})
	})
}


			//cinfo, err := store.GetByCode(info.Code)
			//So(err, ShouldBeNil)
			//So(cinfo.GetUserID(), ShouldEqual, info.UserID)

			err = store.RemoveByCode(info.Code)
			So(err, ShouldBeNil)

			cinfo, err = store.GetByCode(info.Code)
			So(err, ShouldBeNil)
			So(cinfo, ShouldBeNil)
		})

		Convey("Test access token store", func() {
			info := &models.Token{
				ClientID:        "1",
				UserID:          "1_1",
				RedirectURI:     "http://localhost/",
				Scope:           "all",
				Access:          "1_1_1",
				AccessCreateAt:  time.Now(),
				AccessExpiresIn: time.Second * 5,
			}
			err := store.Create(info)
			So(err, ShouldBeNil)

			ainfo, err := store.GetByAccess(info.GetAccess())
			So(err, ShouldBeNil)
			So(ainfo.GetUserID(), ShouldEqual, info.GetUserID())

			err = store.RemoveByAccess(info.GetAccess())
			So(err, ShouldBeNil)

			ainfo, err = store.GetByAccess(info.GetAccess())
			So(err, ShouldBeNil)
			So(ainfo, ShouldBeNil)
		})

		Convey("Test refresh token store", func() {
			info := &models.Token{
				ClientID:         "1",
				UserID:           "1_2",
				RedirectURI:      "http://localhost/",
				Scope:            "all",
				Access:           "1_2_1",
				AccessCreateAt:   time.Now(),
				AccessExpiresIn:  time.Second * 5,
				Refresh:          "1_2_2",
				RefreshCreateAt:  time.Now(),
				RefreshExpiresIn: time.Second * 15,
			}
			err := store.Create(info)
			So(err, ShouldBeNil)

			ainfo, err := store.GetByAccess(info.GetAccess())
			So(err, ShouldBeNil)
			So(ainfo.GetUserID(), ShouldEqual, info.GetUserID())

			err = store.RemoveByAccess(info.GetAccess())
			So(err, ShouldBeNil)

			ainfo, err = store.GetByAccess(info.GetAccess())
			So(err, ShouldBeNil)
			So(ainfo, ShouldBeNil)

			rinfo, err := store.GetByRefresh(info.GetRefresh())
			So(err, ShouldBeNil)
			So(rinfo.GetUserID(), ShouldEqual, info.GetUserID())

			err = store.RemoveByRefresh(info.GetRefresh())
			So(err, ShouldBeNil)

			rinfo, err = store.GetByRefresh(info.GetRefresh())
			So(err, ShouldBeNil)
			So(rinfo, ShouldBeNil)
		})
	})
}
*/

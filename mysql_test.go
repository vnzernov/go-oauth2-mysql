package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/oauth2.v3/models"
	"testing"
	"time"
)

const (
	dsn = "root@tcp(localhost:3306)/myapp_test?charset=utf8"
)

func TestTokenStore(t *testing.T) {
	Convey("Test mysql token store", t, func() {
		store := NewDefaultStore(NewConfig(dsn))
		defer store.clean()

		Convey("Test authorization code store", func() {
			info := &models.Token{
				ClientID:      "1",
				UserID:        "1_1",
				RedirectURI:   "http://localhost/",
				Scope:         "all",
				Code:          "11_11_11",
				CodeCreateAt:  time.Now(),
				CodeExpiresIn: time.Second * 5,
			}
			err := store.Create(info)
			So(err, ShouldBeNil)

			cinfo, err := store.GetByCode(info.Code)
			So(err, ShouldBeNil)
			So(cinfo.GetUserID(), ShouldEqual, info.UserID)

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

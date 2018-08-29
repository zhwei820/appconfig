package group_service_test

import (
	_ "back/appconfig/utils/gotests"

	"net/http"
	"net/http/httptest"
	"testing"
	_ "back/appconfig/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/bitly/go-simplejson"
	"fmt"
	"bytes"
	"back/appconfig/utils/random"
)

func TestCreate(t *testing.T) {
	str := random.RandString(6)
	r, _ := http.NewRequest("POST", fmt.Sprintf("/api/group/group/?"),
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"group": "string%v",
								"groupname": "string%v",
  								"permission": "string%v"
								}`, str, str, str))))

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	var resId int

	Convey("detail\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())
			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)

			resId = j.Get("data").MustInt() // obj id

		})
	})

	testList(t)
	testDetail(t, resId)
	testUpdate(t, resId)

}
func testList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/group/group?page=1&page_size=10", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestList", "code", w.Code, w.Body.String())

	Convey("list\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())

			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)
		})
	})

}

func testDetail(t *testing.T, resId int) {
	r, _ := http.NewRequest("GET", fmt.Sprintf("/api/group/group/%v?", resId), nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	Convey("detail\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())
			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)
		})
	})
}

func testUpdate(t *testing.T, resId int) {
	str := random.RandString(6)

	r, _ := http.NewRequest("PUT", fmt.Sprintf("/api/group/group/%v?", resId),
		bytes.NewBuffer([]byte(fmt.Sprintf(`{"group": "string%v",
								"groupname": "string%v",
  								"permission": "string%v"
								}`, str, str, str))))

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	//beego.Trace("TestDetail", "code", w.Code, w.Body.String())

	Convey("detail\n", t, func() {
		Convey("status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
		Convey("result Code ", func() {
			j, _ := simplejson.NewJson(w.Body.Bytes())
			code := j.Get("code").MustInt()
			So(code, ShouldEqual, 0)
		})
	})
}

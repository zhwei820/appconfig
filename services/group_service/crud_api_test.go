package group_service_test

import (
	_ "back/appconfig/utils/gotests"

	"net/http"
	"net/http/httptest"
	"testing"
	_ "back/appconfig/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
)

// TestGet is a sample to run an endpoint test
func TestList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/group/group?page=-1&page_size=10", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code", w.Code, w.Body.String())

	Convey("List\n", t, func() {
		Convey("Status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

func TestDetail(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/group/group/1?", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code", w.Code, w.Body.String())

	Convey("List\n", t, func() {
		Convey("Status Code ", func() {
			So(w.Code, ShouldEqual, 200)
		})
	})
}

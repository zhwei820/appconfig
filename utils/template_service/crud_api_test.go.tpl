package {{ model_data.obj }}_service_test

import (
	_ "back/natsmicro/utils/gotests"

	"net/http"
	"net/http/httptest"
	"testing"
	_ "back/natsmicro/routers"
	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/bitly/go-simplejson"
	"fmt"
	"bytes"
	"back/natsmicro/utils/random"
	"io"
)

func getRandomInput() io.Reader {
	str := random.RandNumString(6)

	return bytes.NewBuffer([]byte(fmt.Sprintf(`{
	                            {{ model_data.testInput }}
								}`,
								{% for item in model_data.data %} str, {% endfor %}
								)))
}

func TestCreate(t *testing.T) {
	r, _ := http.NewRequest("POST", fmt.Sprintf("/api/{{ model_data.obj }}/{{ model_data.obj }}/?"),
		getRandomInput())

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

			resId = j.Get("data").MustInt() // model_data.obj id

		})
	})

	testList(t)
	testDetail(t, resId)
	testUpdate(t, resId)

}
func testList(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/{{ model_data.obj }}/{{ model_data.obj }}?page=1&page_size=10", nil)
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
	r, _ := http.NewRequest("GET", fmt.Sprintf("/api/{{ model_data.obj }}/{{ model_data.obj }}/%v?", resId), nil)
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

	r, _ := http.NewRequest("PUT", fmt.Sprintf("/api/{{ model_data.obj }}/{{ model_data.obj }}/%v?", resId),
		getRandomInput())

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

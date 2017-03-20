package cookies

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var cookieMng = New("user", &Conf{})

func TestCookieGetSet(t *testing.T) {
	var (
		w = httptest.NewRecorder()
		m = map[string]string{"username": "wind85"}
	)

	cookieMng.Set(w, nil, m)

	r := &http.Request{Header: http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}}

	vals := cookieMng.Get(w, r)
	if vals["username"] != m["username"] {
		t.Fatal("Something went wrong\n")
	} else {
		t.Logf("usename is wind85 as expected: %s\n", vals["username"])
	}
}

func TestDelCookie(t *testing.T) {

	var (
		w = httptest.NewRecorder()
		m = map[string]string{"username": "wind85"}
	)

	cookieMng.Set(w, nil, m)
	cookieMng.Del(w, nil)

}

// package cookies provide a secure implementation based on gorilla securecookie
// to manage cookies, provides 3 functions SetCookieVal GetCookieVal and DelCookie
// The Cookie data type is an implementation the the CookieMng ( cookie manager )
// With this simple type is possible to exchange cookies securely over a non secure
// connection values are encrypted and decripted back, I don't suggest to pass
// sensitive data unless with an other encryption layer on top.
// It should be noted that for testing porpoises the HttpOnly and Secure are set
// to false, you must change that when using it on a production environment,to
// true.

package cookies

import (
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

type CookieMng interface {
	SetCookieVal(w http.ResponseWriter, r *http.Request, val map[string]string)
	GetCookieVal(w http.ResponseWriter, r *http.Request) map[string]string
	DelCookie(w http.ResponseWriter, r *http.Request)
}

type Cookies struct {
	secure *securecookie.SecureCookie
	conf   *Conf
	name   string
}

type Conf struct {
	HttpOnly bool
	Secure   bool
	MaxAge   int
}

func New(name string, conf *Conf) *Cookies {
	return &Cookies{
		securecookie.New(
			securecookie.GenerateRandomKey(64),
			securecookie.GenerateRandomKey(32),
		),
		name,
		conf,
	}
}

func (c *Cookies) SetCookieVal(w http.ResponseWriter, r *http.Request, val map[string]string) {
	c.setCookie(w, r, val, false, false, 0, time.Now().Add(168*time.Hour))
}

func (c *Cookies) GetCookieVal(w http.ResponseWriter, r *http.Request) map[string]string {
	value := make(map[string]string)

	if cookie, err := r.Cookie(c.name); err == nil {
		if err = c.secure.Decode(c.name, cookie.Value, &value); err != nil {
			http.Error(w, "Ops internal server error\n", http.StatusInternalServerError)
			return nil
		}
	}

	return value
}

func (c *Cookies) DelCookie(w http.ResponseWriter, r *http.Request) {
	c.setCookie(w, r, nil, false, false, -1, time.Now().Add(168*time.Hour))
}

func (c *Cookies) setCookie(w http.ResponseWriter, r *http.Request,
	val map[string]string, secure, httponly bool, age int, expiration time.Time) {

	if encoded, err := c.secure.Encode(c.name, val); err == nil {
		cookie := &http.Cookie{
			Name:     c.name,
			Value:    encoded,
			Path:     "/",
			HttpOnly: httponly,
			Secure:   secure,
			MaxAge:   age,
			Expires:  expiration,
		}

		http.SetCookie(w, cookie)
	}
}

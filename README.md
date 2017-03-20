# Cookies

[![GoDoc](https://godoc.org/github.com/wind85/cookies?status.svg)](https://godoc.org/github.com/wind85/cookies)
[![Build Status](https://travis-ci.org/wind85/cookies.svg?branch=master)](https://travis-ci.org/wind85/cookies)
[![Coverage Status](https://coveralls.io/repos/github/wind85/cookies/badge.svg?branch=master)](https://coveralls.io/github/wind85/cookies?branch=master)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
### Cookies package
This is a small package the provides a layer of abstraction over gorilla secure cookie, in
order to give a plug and play usage. Security is a paramount therefore cookies are encrypted
only with secure randomly generate keys. The package provides only the following 3 methods:

- New accept two value the cookie name and a configuration ( the Conf struct ).
- Set method accept a map of string keys and string values and adds created and adds the cookie.
- Get retrieves the cookie map for the request and return a map of string keys and string values.

### How to use it

Pretty simple, there is only one method to create a new parser just call 
```
  cookiemng := cookies.New("cookiename",&Conf{}) 
  // default conf uses default values not secure, max age 0 and not httponly.
```
To put, retrieve and delete cookies to like so:
```
  cookiemng.Set(w, r, map[string]string{"testkey":"testvalue"}) 
 // sets the cookie with the following map
  values := cookiemng.Get(w,r)
 // values containes all the value previously set into the cookie
 cookiemng.Del(w,r)
 // clears the cookie from the request
```
That's pretty much it.

#### Philosophy
This software is developed following the "mantra" keep it simple, stupid or better known as
KISS. Something so simple like a cache with auto eviction should not required over engineered 
solutions. Though it provides most of the functionality needed by generic configuration files, 
and most important of all meaning full error messages.

#### Disclaimer
This software in alpha quality, don't use it in a production environment, it's not even completed.

#### Thank You Notes
None.

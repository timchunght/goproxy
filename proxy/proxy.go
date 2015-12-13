package goproxy

import (
	"errors"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/unrolled/render"
	"fmt"
)

var (
	ErrInvalidService = errors.New("invalid service/version")
	r = render.New()

)

func Proxy_pass(scheme string, host string, w http.ResponseWriter, req *http.Request) error {

	// Try to connect
	_, err := net.Dial("tcp", host)
	if err != nil {

		return errors.New(fmt.Sprintf("Connection to %s failed", host))
	}
	httputil.NewSingleHostReverseProxy(&url.URL{Scheme: scheme, Host: host}).ServeHTTP(w, req)
	return nil
}

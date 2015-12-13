package main

import (
	"goproxy/proxy"
	"net/http"
	"os"
	"log"
	"github.com/unrolled/render"
)

var (

	r = render.New()
)
func main() {
	// http.HandleFunc("/", goproxy.NewMultipleHostReverseProxy(ServiceRegistry))
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		if req.URL.String() == "/upload" {
			err := goproxy.Proxy_pass("http", "159.203.84.115:8080", w, req)
			if err != nil {
				r.JSON(w, http.StatusUnauthorized, map[string]string{"message": "please use /upload for upload"})
				log.Println(err)
			}
		} else {

			r.JSON(w, http.StatusUnauthorized, map[string]string{"message": "please use /upload for upload"})

		}
	
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

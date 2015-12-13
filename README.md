# goproxy

goproxy is a Reverse Proxy written in Go that can be used to communicate with database before routing request to backend inspired by proxy_machine

# Example: 

In this example, I am making an upload server. Say I have a backend server at ``localhost:8080`` (call it upstream) and by using ``Proxy_pass``, I am redirecting the request to ``localhost:8080``. 

In an actual application, you might have a database that stores the file location with its corresponding ``host`` info, all you have to do is use that ``host`` as parameter to ``Proxy_pass`` and you are done. 

This project aims to solve some features that nginx does not have like retrieving information from database. Using Go, we can have a similar performance without losing the ability to communicate with our database. Database communication is the only advantage a custom reverse proxy has over nginx. In fact, the upstream server uses nginx.

```go
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
			err := goproxy.Proxy_pass("http", "localhost:8080", w, req)
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

```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "static")
	e.GET("/", func(c echo.Context) (err error) {
		pusher, ok := c.Response().Writer.(http.Pusher)
		if ok {
			if err = pusher.Push("/app.css", nil); err != nil {
				return
			}
			if err = pusher.Push("/app.js", nil); err != nil {
				return
			}
			if err = pusher.Push("/echo.png", nil); err != nil {
				return
			}
		}
		return c.File("index.html")
	})
	e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
}

/*
Step 1: Generate a self-signed X.509 TLS certificate
		Run the following command to generate cert.pem and key.pem files:
		go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost

Step 2: Register a route to serve web assets
		e.Static("/", "static")

Step 3: Create a handler to serve index.html and push it’s dependencies
		e.GET("/", func(c echo.Context) (err error) {
		  pusher, ok := c.Response().Writer.(http.Pusher)
		  if ok {
		    if err = pusher.Push("/app.css", nil); err != nil {
		      return
		    }
		    if err = pusher.Push("/app.js", nil); err != nil {
		      return
		    }
		    if err = pusher.Push("/echo.png", nil); err != nil {
		      return
		    }
		  }
		  return c.File("index.html")
		})

Step 4: Configure TLS server using cert.pem and key.pem
		e.StartTLS(":1323", "cert.pem", "key.pem")

Step 5: Start the server and browse to https://localhost:1323
		Protocol: HTTP/2.0
		Host: localhost:1323
		Remote Address: [::1]:60288
		Method: GET
		Path: /
*/

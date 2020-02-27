package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var index = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Upstream Server</title>
		<style>
			h1, p {
				font-weight: 300;
			}
		</style>
	</head>
	<body>
		<p>
			Hello from upstream server %s
		</p>
	</body>
	</html>
`

func main() {
	name := os.Args[1]
	port := os.Args[2]
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, fmt.Sprintf(index, name))
	})
	e.Logger.Fatal(e.Start(port))
}

/*
How to setup Nginx proxy server with Echo?

Step 1: Install Nginx
https://www.nginx.com/resources/wiki/start/topics/tutorials/install

Step 2: Configure Nginx
Create a file /etc/nginx/sites-enabled/localhost with the following content:

upstream localhost {
  server localhost:8081;
  server localhost:8082;
}

server {
  listen          8080;
  server_name     localhost;
  access_log      /var/log/nginx/localhost.access.log combined;

  location / {
    proxy_pass      http://localhost;
  }
}

Step 3: Restart Nginx
service nginx restart

Step 4: Start upstream servers
cd upstream
go run server.go server1 :8081
go run server.go server2 :8082

Step 5: Browse to https://localhost:8080


How to setup Armor proxy server with Echo?

Step 1: Install Armor
https://armor.labstack.com/guide

Step 2: Configure Armor
Create a file /etc/armor/config.yaml with the following content:

address: ":8080"
plugins:
- name: logger
hosts:
  localhost:1323:
    paths:
      "/":
        plugins:
        - name: proxy
          targets:
          - url: http://localhost:8081
          - url: http://localhost:8082
Step 3: Start Armor
armor -c /etc/armor/config.yaml

Change address and hosts per your need.

Step 4 & 5: Follow Nginx recipe
*/

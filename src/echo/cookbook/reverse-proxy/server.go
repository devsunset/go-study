package main

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Setup proxy
	url1, err := url.Parse("http://localhost:8081")
	if err != nil {
		e.Logger.Fatal(err)
	}
	url2, err := url.Parse("http://localhost:8082")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
		{
			URL: url2,
		},
	}
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start(":1323"))
}

/*
Step 1: Identify upstream target URLs
	url1, err := url.Parse("http://localhost:8081")
	if err != nil {
	  e.Logger.Fatal(err)
	}
	url2, err := url.Parse("http://localhost:8082")
	if err != nil {
	  e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
	  {
	    URL: url1,
	  },
	  {
	    URL: url2,
	  },
	}

Step 2: Setup proxy middleware with upstream targets
	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	g := e.Group("/blog")
	g.Use(middleware.Proxy(...))

Step 3: Start upstream servers
	cd upstream
	go run server.go server1 :8081
	go run server.go server2 :8082

Step 3: Start the proxy server
	go run server.go

Step 4: Browse to http://localhost:1323
*/

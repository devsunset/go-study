package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	//http://localhost:1323
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
		// #Request
		// Context#Bind(i interface{})
		// Echo#Binder
		// Context#FormValue(name string)
		// Echo#BindUnmarshaler
		// Context#QueryParam(name string)
		// Context#Param(name string) string
		// Echo#Validator (https://github.com/go-playground/validator)

		// #Response
		// Context#String(code int, s string)
		// Context#HTML(code int, html string)
		// Context#HTMLBlob(code int, b []byte)
		// Context#JSON(code int, i interface{})
		// Context#JSON()
		// Context#JSONPretty(code int, i interface{}, indent string)
		// Context#JSONBlob(code int, b []byte)
		// Context#JSONP(code int, callback string, i interface{})
		// Context#XML(code int, i interface{})
		// Context#XML()
		// Context#XMLPretty(code int, i interface{}, indent string)
		// Context#XMLBlob(code int, b []byte)
		// Context#File(file string)
		// Context#Attachment(file, name string)
		// Context#Inline(file, name string)
		// Context#Blob(code int, contentType string, b []byte)
		// Context#Stream(code int, contentType string, r io.Reader)
		// Context#NoContent(code int)
		// Context#Redirect(code int, url string)
		// Context#Redirect(code int, url string)
		// Context#Response#After(func())
	})

	// #Handling Request
	// Bind json, xml, form or query payload into Go struct based on Content-Type request header.
	// Render response as json or xml with status code.
	type User struct {
		Name  string `json:"name" xml:"name" form:"name" query:"name"`
		Email string `json:"email" xml:"email" form:"email" query:"email"`
	}

	e.GET("/members", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})

	// #Routing
	//  e.GET("/users/:id", getUser)
	//	e.POST("/users", saveUser)
	//	e.PUT("/users/:id", updateUser)
	//	e.DELETE("/users/:id", deleteUser)

	// Echo.Any(path string, h Handler)
	// Echo.Match(methods []string, path string, h Handler)
	// Echo#Group(prefix string, m ...Middleware) *Group
	// Echo#URI(handler HandlerFunc, params ...interface{})
	// Echo#Reverse(name string, params ...interface{})

	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/saveFile", saveFile)

	e.GET("/writeCookie", writeCookie)
	e.GET("/readCookie", readCookie)
	e.GET("/readAllCookies", readAllCookies)

	fmt.Println(e.Routes())
	// Echo#Routes() []*Route
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("static/routes.json", data, 0644)

	// #Static Content
	e.Static("/static", "static")

	// Echo#Static(prefix, root string)
	// Usage 1
	// e.Static("/static", "assets")   /static/* /static/js/main.js -> assets/js/main.js
	// Usage 2
	// e.Static("/", "assets")         /*        /js/main.js -> assets/js/main.js

	// Echo#File(path, file string)
	// Usage 1
	// e.File("/", "public/index.html") 		     public/index.html
	// Usage 2
	// e.File("/favicon.ico", "images/favicon.ico")  images/favicon.ico

	// #Middleware
	////////////////////////////////////////////////////
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /rlm")
			return next(c)
		}
	}
	//http://localhost:1323/rlm
	e.GET("/rlm", func(c echo.Context) error {
		return c.String(http.StatusOK, "/rlm")
	}, track)
	////////////////////////////////////////////////////

	// #Templates

	t := &Template{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}

	e.Renderer = t
	e.GET("/hello", Hello)

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}
	e.Renderer = renderer

	// Named route "foobar"
	e.GET("/something", func(c echo.Context) error {
		return c.Render(http.StatusOK, "something.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	// #Custom Error
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Logger.Fatal(e.Start(":1323"))
}

// Path Parameters (Param)
// http://localhost:1323/users/Joe
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// Query Parameters (QueryParam)
// http://localhost:1323/show?team=x-men&member=wolverine
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// Form application/x-www-form-urlencoded (FormValue)
// curl -F "name=Joe Smith" -F "email=joe@labstack.com" http://localhost:1323/save
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// Form multipart/form-data (FormValue + FormFile)
// curl -F "name=Joe Smith" -F "avatar=@/path/to/your/avatar.png" http://localhost:1323/save
func saveFile(c echo.Context) error {
	// Get name
	name := c.FormValue("name")
	// Get avatar
	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// Source
	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"</b>")
}

// Create a Cookie
func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "username"
	cookie.Value = "jon"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

// Read a Cookie
func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("username")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Name)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "read a cookie")
}

// Read all the Cookies
func readAllCookies(c echo.Context) error {
	for _, cookie := range c.Cookies() {
		fmt.Println(cookie.Name)
		fmt.Println(cookie.Value)
	}
	return c.String(http.StatusOK, "read all the cookies")
}

// Error Pages
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "hello", "World")
}

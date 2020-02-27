package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//  #Root Level (Before router)
//  Echo#Pre()
//	HTTPSRedirect
//	HTTPSWWWRedirect
//	WWWRedirect
//	NonWWWRedirect
//	AddTrailingSlash
//	RemoveTrailingSlash
//	MethodOverride
//	Rewrite

//  #Root Level (After router)
//  Echo#Use()
//	BodyLimit
//	Logger
//	Gzip
//	Recover
//	BasicAuth
//	JWTAuth
//	Secure
//	CORS
//	Static

//  #Group Level
// 	e := echo.New()
//	admin := e.Group("/admin", middleware.BasicAuth())

//	#Route Level
//	e := echo.New()
//	e.GET("/", <Handler>, <Middleware...>)

//  #Skipping Middleware
//	Skipper func(c echo.Context) bool
//	e := echo.New()
//	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
//		Skipper: func(c echo.Context) bool {
//			if strings.HasPrefix(c.Request().Host, "localhost") {
//				return true
//			}
//			return false
//		},
//	}))

func main() {
	// Echo instance
	e := echo.New()

	// #Basic Auth Middleware
	// Basic auth middleware provides an HTTP basic authentication.

	/*
		e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
			// Be careful to use constant time comparison to prevent timing attacks

			if username == "joe" && password == "secret" {
				return true, nil
			}
			//		if subtle.ConstantTimeCompare([]byte(username), []byte("joe")) == 1 &&
			//			subtle.ConstantTimeCompare([]byte(password), []byte("secret")) == 1 {
			//			return true, nil
			//		}

			return false, nil
		}))
	*/

	/*
			// Custom Configuration
		 	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{}))

			//Configuration
			BasicAuthConfig struct {
			  // Skipper defines a function to skip middleware.
			  Skipper Skipper

			  // Validator is a function to validate BasicAuth credentials.
			  // Required.
			  Validator BasicAuthValidator

			  // Realm is a string to define realm attribute of BasicAuth.
			  // Default value "Restricted".
			  Realm string
			}

			//Default Configuration
			DefaultBasicAuthConfig = BasicAuthConfig{
				Skipper: DefaultSkipper,
			}
	*/

	// #Body Dump Middleware
	//	Body dump middleware captures the request and response payload and calls the registered handler.
	//	Generally used for debugging/logging purpose. Avoid using it
	//	if your request/response payload is huge e.g. file upload/download,
	//	but if you still need to, add an exception for your endpoints in the skipper function.

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Println("Logging...")
	}))

	/*
			// Custom Configuration
		 	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{}))

			//Configuration
			BodyDumpConfig struct {
			  // Skipper defines a function to skip middleware.
			  Skipper Skipper

			  // Handler receives request and response payload.
			  // Required.
			  Handler BodyDumpHandler
			}

			//Default Configuration
			DefaultBodyDumpConfig = BodyDumpConfig{
			  Skipper: DefaultSkipper,
			}
	*/

	// #Body Limit Middleware
	// Body limit middleware sets the maximum allowed size for a request body

	e.Use(middleware.BodyLimit("2M"))

	/*
		// Custom Configuration
		e.Use(middleware.BodyLimitWithConfig(middleware.BodyLimitConfig{}))

		//Configuration
		BodyLimitConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Maximum allowed size for a request body, it can be specified
		  // as `4x` or `4xB`, where x is one of the multiple from K, M, G, T or P.
		  Limit string `json:"limit"`
		}

		//Default Configuration
		DefaultBodyLimitConfig = BodyLimitConfig{
		  Skipper: DefaultSkipper,
		}
	*/

	// #Casbin Auth Middleware
	//  Casbin is a powerful and efficient open-source access control library for Go
	//	ACL (Access Control List)
	//	ACL with superuser
	//	ACL without users: especially useful for systems that don’t have authentication or user log-ins.
	//	ACL without resources: some scenarios may target for a type of resources instead of an individual resource by using permissions like write-article, read-log. It doesn’t control the access to a specific article or log.
	//	RBAC (Role-Based Access Control)
	//	RBAC with resource roles: both users and resources can have roles (or groups) at the same time.
	//	RBAC with domains/tenants: users can have different role sets for different domains/tenants.
	//	ABAC (Attribute-Based Access Control)
	//	RESTful
	//	Deny-override: both allow and deny authorizations are supported, deny overrides the allow.

	/*
		Dependencies
		import (
		  "github.com/casbin/casbin"
		  casbin_mw "github.com/labstack/echo-contrib/casbin"
		)

		enforcer, err := casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")
		e.Use(casbin_mw.Middleware(enforcer))

		//  Custom Configuration
		ce := casbin.NewEnforcer("casbin_auth_model.conf", "")
		ce.AddRoleForUser("alice", "admin")
		ce.AddPolicy(...)
		e.Use(casbin_mw.MiddlewareWithConfig(casbin_mw.Config{
		  Enforcer: ce,
		}))

		//Configuration
		// Config defines the config for CasbinAuth middleware.
		Config struct {
		  // Skipper defines a function to skip middleware.
		  Skipper middleware.Skipper

		  // Enforcer CasbinAuth main rule.
		  // Required.
		  Enforcer *casbin.Enforcer
		}

		//Default Configuration
		// DefaultConfig is the default CasbinAuth middleware config.
			DefaultConfig = Config{
			  Skipper: middleware.DefaultSkipper,
			}
	*/

	// #CORS Middleware
	// CORS middleware implements CORS specification. CORS gives web servers cross-domain access controls, which enable secure cross-domain data transfers.

	e.Use(middleware.CORS())

	/*
		// Custom Configuration
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		  AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		  AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))

		//Configuration
		CORSConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // AllowOrigin defines a list of origins that may access the resource.
		  // Optional. Default value []string{"*"}.
		  AllowOrigins []string `json:"allow_origins"`

		  // AllowMethods defines a list methods allowed when accessing the resource.
		  // This is used in response to a preflight request.
		  // Optional. Default value DefaultCORSConfig.AllowMethods.
		  AllowMethods []string `json:"allow_methods"`

		  // AllowHeaders defines a list of request headers that can be used when
		  // making the actual request. This in response to a preflight request.
		  // Optional. Default value []string{}.
		  AllowHeaders []string `json:"allow_headers"`

		  // AllowCredentials indicates whether or not the response to the request
		  // can be exposed when the credentials flag is true. When used as part of
		  // a response to a preflight request, this indicates whether or not the
		  // actual request can be made using credentials.
		  // Optional. Default value false.
		  AllowCredentials bool `json:"allow_credentials"`

		  // ExposeHeaders defines a whitelist headers that clients are allowed to
		  // access.
		  // Optional. Default value []string{}.
		  ExposeHeaders []string `json:"expose_headers"`

		  // MaxAge indicates how long (in seconds) the results of a preflight request
		  // can be cached.
		  // Optional. Default value 0.
		  MaxAge int `json:"max_age"`
		}

		//Default Configuration
		DefaultCORSConfig = CORSConfig{
		  Skipper:      DefaultSkipper,
		  AllowOrigins: []string{"*"},
		  AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		}
	*/

	// #CSRF Middleware
	// Cross-site request forgery, also known as one-click attack or session riding and abbreviated as CSRF (sometimes pronounced sea-surf) or XSRF, is a type of malicious exploit of a website where unauthorized commands are transmitted from a user that the website trusts.

	e.Use(middleware.CSRF())

	/*
		// Custom Configuration
		e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		  TokenLookup: "header:X-XSRF-TOKEN",
		}))

		//Configuration
		CSRFConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // TokenLength is the length of the generated token.
		  TokenLength uint8 `json:"token_length"`
		  // Optional. Default value 32.

		  // TokenLookup is a string in the form of "<source>:<key>" that is used
		  // to extract token from the request.
		  // Optional. Default value "header:X-CSRF-Token".
		  // Possible values:
		  // - "header:<name>"
		  // - "form:<name>"
		  // - "query:<name>"
		  TokenLookup string `json:"token_lookup"`

		  // Context key to store generated CSRF token into context.
		  // Optional. Default value "csrf".
		  ContextKey string `json:"context_key"`

		  // Name of the CSRF cookie. This cookie will store CSRF token.
		  // Optional. Default value "csrf".
		  CookieName string `json:"cookie_name"`

		  // Domain of the CSRF cookie.
		  // Optional. Default value none.
		  CookieDomain string `json:"cookie_domain"`

		  // Path of the CSRF cookie.
		  // Optional. Default value none.
		  CookiePath string `json:"cookie_path"`

		  // Max age (in seconds) of the CSRF cookie.
		  // Optional. Default value 86400 (24hr).
		  CookieMaxAge int `json:"cookie_max_age"`

		  // Indicates if CSRF cookie is secure.
		  // Optional. Default value false.
		  CookieSecure bool `json:"cookie_secure"`

		  // Indicates if CSRF cookie is HTTP only.
		  // Optional. Default value false.
		  CookieHTTPOnly bool `json:"cookie_http_only"`
		}

		//Default Configuration
		DefaultCSRFConfig = CSRFConfig{
		  Skipper:      DefaultSkipper,
		  TokenLength:  32,
		  TokenLookup:  "header:" + echo.HeaderXCSRFToken,
		  ContextKey:   "csrf",
		  CookieName:   "_csrf",
		  CookieMaxAge: 86400,
		}
	*/

	// #Gzip Middleware
	// Gzip middleware compresses HTTP response using gzip compression scheme.

	e.Use(middleware.Gzip())

	/*
		// Custom Configuration
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		  Level: 5,
		}))

		//Configuration
		GzipConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Gzip compression level.
		  // Optional. Default value -1.
		  Level int `json:"level"`
		}

		//Default Configuration
		DefaultGzipConfig = GzipConfig{
		  Skipper: DefaultSkipper,
		  Level:   -1,
		}
	*/

	// #Jaeger Tracing Middleware
	// Trace requests on Echo framework with Jaeger Tracing Middleware.

	/*
		 //Dependencies
			import (
			    "github.com/labstack/echo-contrib/jaegertracing"
			    "github.com/labstack/echo/v4"
			)


		 //Enable tracing middleware
		  c := jaegertracing.New(e, nil)
		  defer c.Close()

		// Custom Configuration
		JAEGER_AGENT_HOST=192.168.1.10 JAEGER_AGENT_PORT=6831 ./myserver

		Property	Description
		JAEGER_SERVICE_NAME	The service name
		JAEGER_AGENT_HOST	The hostname for communicating with agent via UDP
		JAEGER_AGENT_PORT	The port for communicating with agent via UDP
		JAEGER_ENDPOINT	The HTTP endpoint for sending spans directly to a collector, i.e. http://jaeger-collector:14268/api/traces
		JAEGER_USER	Username to send as part of “Basic” authentication to the collector endpoint
		JAEGER_PASSWORD	Password to send as part of “Basic” authentication to the collector endpoint
		JAEGER_REPORTER_LOG_SPANS	Whether the reporter should also log the spans
		JAEGER_REPORTER_MAX_QUEUE_SIZE	The reporter’s maximum queue size
		JAEGER_REPORTER_FLUSH_INTERVAL	The reporter’s flush interval, with units, e.g. “500ms” or “2s” ([valid units][timeunits])
		JAEGER_SAMPLER_TYPE	The sampler type
		JAEGER_SAMPLER_PARAM	The sampler parameter (number)
		JAEGER_SAMPLER_MANAGER_HOST_PORT	The HTTP endpoint when using the remote sampler, i.e. http://jaeger-agent:5778/sampling
		JAEGER_SAMPLER_MAX_OPERATIONS	The maximum number of operations that the sampler will keep track of
		JAEGER_SAMPLER_REFRESH_INTERVAL	How often the remotely controlled sampler will poll jaeger-agent for the appropriate sampling strategy, with units, e.g. “1m” or “30s” ([valid units][timeunits])
		JAEGER_TAGS	A comma separated list of name = value tracer level tags, which get added to all reported spans. The value can also refer to an environment variable using the format ${envVarName:default}, where the :default is optional, and identifies a value to be used if the environment variable cannot be found
		JAEGER_DISABLED	Whether the tracer is disabled or not. If true, the default opentracing.NoopTracer is used.
		JAEGER_RPC_METRICS	Whether to store RPC metrics

		//Middleware Skipper
		//A middleware skipper can be passed to avoid tracing spans to certain URLs
		package main
		import (
			"strings"
		    "github.com/labstack/echo-contrib/jaegertracing"
		    "github.com/labstack/echo/v4"
		)

		// urlSkipper ignores metrics route on some middleware
		func urlSkipper(c echo.Context) bool {
		    if strings.HasPrefix(c.Path(), "/testurl") {
		        return true
		    }
		    return false
		}

		func main() {
		    e := echo.New()
		    // Enable tracing middleware
		    c := jaegertracing.New(e, urlSkipper)
		    defer c.Close()

		    e.Logger.Fatal(e.Start(":1323"))
		}

		//TraceFunction
		//This is a wrapper function that can be used to seamlessly add a span for the duration of the invoked function. There is no need to change function arguments.
		package main
		import (
		    "github.com/labstack/echo-contrib/jaegertracing"
		    "github.com/labstack/echo/v4"
		    "net/http"
		    "time"
		)
		func main() {
		    e := echo.New()
		    // Enable tracing middleware
		    c := jaegertracing.New(e, nil)
		    defer c.Close()
		    e.GET("/", func(c echo.Context) error {
		        // Wrap slowFunc on a new span to trace it's execution passing the function arguments
				jaegertracing.TraceFunction(c, slowFunc, "Test String")
		        return c.String(http.StatusOK, "Hello, World!")
		    })
		    e.Logger.Fatal(e.Start(":1323"))
		}

		// A function to be wrapped. No need to change it's arguments due to tracing
		func slowFunc(s string) {
			time.Sleep(200 * time.Millisecond)
			return
		}

		//CreateChildSpan
		//For more control over the Span, the function CreateChildSpan can be called giving control on data to be appended to the span like log messages, baggages and tags.
		package main
		import (
		    "github.com/labstack/echo-contrib/jaegertracing"
		    "github.com/labstack/echo/v4"
		)
		func main() {
		    e := echo.New()
		    // Enable tracing middleware
		    c := jaegertracing.New(e, nil)
		    defer c.Close()
		    e.GET("/", func(c echo.Context) error {
		        // Do something before creating the child span
		        time.Sleep(40 * time.Millisecond)
		        sp := jaegertracing.CreateChildSpan(c, "Child span for additional processing")
		        defer sp.Finish()
		        sp.LogEvent("Test log")
		        sp.SetBaggageItem("Test baggage", "baggage")
		        sp.SetTag("Test tag", "New Tag")
		        time.Sleep(100 * time.Millisecond)
		        return c.String(http.StatusOK, "Hello, World!")
		    })
		    e.Logger.Fatal(e.Start(":1323"))
		}
	*/

	// #JWT Middleware
	// JWT provides a JSON Web Token (JWT) authentication middleware.

	// e.Use(middleware.JWT([]byte("secret")))

	/*
		// Custom Configuration
		e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		  SigningKey: []byte("secret"),
		  TokenLookup: "query:token",
		}))

		//Configuration
		JWTConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Signing key to validate token.
		  // Required.
		  SigningKey interface{}

		  // Signing method, used to check token signing method.
		  // Optional. Default value HS256.
		  SigningMethod string

		  // Context key to store user information from the token into context.
		  // Optional. Default value "user".
		  ContextKey string

		  // Claims are extendable claims data defining token content.
		  // Optional. Default value jwt.MapClaims
		  Claims jwt.Claims

		  // TokenLookup is a string in the form of "<source>:<name>" that is used
		  // to extract token from the request.
		  // Optional. Default value "header:Authorization".
		  // Possible values:
		  // - "header:<name>"
		  // - "query:<name>"
		  // - "cookie:<name>"
		  TokenLookup string

		  // AuthScheme to be used in the Authorization header.
		  // Optional. Default value "Bearer".
		  AuthScheme string
		}

		//Default Configuration
		DefaultJWTConfig = JWTConfig{
		  Skipper:       DefaultSkipper,
		  SigningMethod: AlgorithmHS256,
		  ContextKey:    "user",
		  TokenLookup:   "header:" + echo.HeaderAuthorization,
		  AuthScheme:    "Bearer",
		  Claims:        jwt.MapClaims{},
		}
	*/

	// #Key Auth Middleware
	// Key auth middleware provides a key based authentication.

	//	e.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
	//		return key == "valid-key", nil
	//	}))

	/*
		// Custom Configuration
		e := echo.New()
		e.Use(middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		  KeyLookup: "query:api-key",
		}))

		//Configuration
		KeyAuthConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // KeyLookup is a string in the form of "<source>:<name>" that is used
		  // to extract key from the request.
		  // Optional. Default value "header:Authorization".
		  // Possible values:
		  // - "header:<name>"
		  // - "query:<name>"
		  KeyLookup string `json:"key_lookup"`

		  // AuthScheme to be used in the Authorization header.
		  // Optional. Default value "Bearer".
		  AuthScheme string

		  // Validator is a function to validate key.
		  // Required.
		  Validator KeyAuthValidator
		}

		//Default Configuration
		DefaultKeyAuthConfig = KeyAuthConfig{
		  Skipper:    DefaultSkipper,
		  KeyLookup:  "header:" + echo.HeaderAuthorization,
		  AuthScheme: "Bearer",
		}
	*/

	// #Logger Middleware
	// Logger middleware logs the information about each HTTP request.

	e.Use(middleware.Logger())

	/*
		// Custom Configuration
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		  Format: "method=${method}, uri=${uri}, status=${status}\n",
		}))

		//Configuration
		LoggerConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Tags to constructed the logger format.
		  //
		  // - time_unix
		  // - time_unix_nano
		  // - time_rfc3339
		  // - time_rfc3339_nano
		  // - id (Request ID)
		  // - remote_ip
		  // - uri
		  // - host
		  // - method
		  // - path
		  // - referer
		  // - user_agent
		  // - status
		  // - error
		  // - latency (In nanoseconds)
		  // - latency_human (Human readable)
		  // - bytes_in (Bytes received)
		  // - bytes_out (Bytes sent)
		  // - header:<NAME>
		  // - query:<NAME>
		  // - form:<NAME>
		  // - cookie:<NAME>
		  //
		  // Example "${remote_ip} ${status}"
		  //
		  // Optional. Default value DefaultLoggerConfig.Format.
		  Format string `json:"format"`

		  // Output is a writer where logs are written.
		  // Optional. Default value os.Stdout.
		  Output io.Writer
		}

		//Default Configuration
		DefaultLoggerConfig = LoggerConfig{
		  Skipper: DefaultSkipper,
		  Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
		    `"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
		    `"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
		    `"bytes_out":${bytes_out}}` + "\n",
		  Output: os.Stdout
		}
	*/

	// #Method Override Middleware
	// Method override middleware checks for the overridden method from the request and uses it instead of the original method.

	e.Pre(middleware.MethodOverride())

	/*
		// Custom Configuration
		e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		  Getter: middleware.MethodFromForm("_method"),
		}))

		//Configuration
		MethodOverrideConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Getter is a function that gets overridden method from the request.
		  // Optional. Default values MethodFromHeader(echo.HeaderXHTTPMethodOverride).
		  Getter MethodOverrideGetter
		}

		//Default Configuration
		DefaultMethodOverrideConfig = MethodOverrideConfig{
		  Skipper: DefaultSkipper,
		  Getter:  MethodFromHeader(echo.HeaderXHTTPMethodOverride),
		}
	*/

	// #Prometheus Middleware
	// Prometheus middleware generates metrics for HTTP requests.

	/*
		package main
		import (
		    "github.com/labstack/echo/v4"
		    "github.com/labstack/echo-contrib/prometheus"
		)
		func main() {
		    e := echo.New()
		    // Enable metrics middleware
		    p := prometheus.NewPrometheus("echo", nil)
		    p.Use(e)

		    e.Logger.Fatal(e.Start(":1323"))
		}

		// Custom Configuration
		package main
		import (
		    "github.com/labstack/echo/v4"
		    "github.com/labstack/echo-contrib/prometheus"
		)

		// urlSkipper ignores metrics route on some middleware
		func urlSkipper(c echo.Context) bool {
			if strings.HasPrefix(c.Path(), "/testurl") {
				return true
			}
			return false
		}

		func main() {
		    e := echo.New()
		    // Enable metrics middleware
		    p := prometheus.NewPrometheus("echo", urlSkipper)
		    p.Use(e)

		    e.Logger.Fatal(e.Start(":1323"))
		}
	*/

	// #Proxy Middleware
	// Proxy provides an HTTP/WebSocket reverse proxy middleware. It forwards a request to upstream server using a configured load balancing technique.

	//	url1, err := url.Parse("http://localhost:8081")
	//	if err != nil {
	//	  e.Logger.Fatal(err)
	//	}
	//	url2, err := url.Parse("http://localhost:8082")
	//	if err != nil {
	//	  e.Logger.Fatal(err)
	//	}
	//	e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer([]*middleware.ProxyTarget{
	//	  {
	//	    URL: url1,
	//	  },
	//	  {
	//	    URL: url2,
	//	  },
	//	})))

	/*
		// Custom Configuration
		e.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{}))

		//Configuration
		// ProxyConfig defines the config for Proxy middleware.
		ProxyConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Balancer defines a load balancing technique.
		  // Required.
		  // Possible values:
		  // - RandomBalancer
		  // - RoundRobinBalancer
		  Balancer ProxyBalancer
		}

		//Default Configuration
		Name	Value
		Skipper	DefaultSkipper
	*/

	// #Recover Middleware
	// Recover middleware recovers from panics anywhere in the chain, prints stack trace and handles the control to the centralized HTTPErrorHandler.

	e.Use(middleware.Recover())

	/*
		// Custom Configuration
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		  StackSize:  1 << 10, // 1 KB
		}))

		//Configuration
		RecoverConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Size of the stack to be printed.
		  // Optional. Default value 4KB.
		  StackSize int `json:"stack_size"`

		  // DisableStackAll disables formatting stack traces of all other goroutines
		  // into buffer after the trace for the current goroutine.
		  // Optional. Default value false.
		  DisableStackAll bool `json:"disable_stack_all"`

		  // DisablePrintStack disables printing stack trace.
		  // Optional. Default value as false.
		  DisablePrintStack bool `json:"disable_print_stack"`
		}

		//Default Configuration
		DefaultRecoverConfig = RecoverConfig{
		  Skipper:           DefaultSkipper,
		  StackSize:         4 << 10, // 4 KB
		  DisableStackAll:   false,
		  DisablePrintStack: false,
		}
	*/

	// #Redirect Middleware

	// HTTPS Redirect
	// HTTPS redirect middleware redirects http requests to https. For example, http://labstack.com will be redirected to https://labstack.com.
	// e.Pre(middleware.HTTPSRedirect())

	// HTTPS WWW Redirect
	// HTTPS WWW redirect redirects http requests to www https. For example, http://labstack.com will be redirected to https://www.labstack.com.
	// e.Pre(middleware.HTTPSWWWRedirect())

	// HTTPS NonWWW Redirect
	// HTTPS NonWWW redirect redirects http requests to https non www. For example, http://www.labstack.com will be redirect to https://labstack.com.
	// e.Pre(middleware.HTTPSNonWWWRedirect())

	// WWW Redirect
	// WWW redirect redirects non www requests to www.
	// For example, http://labstack.com will be redirected to http://www.labstack.com.
	// e.Pre(middleware.WWWRedirect())

	// NonWWW Redirect
	// NonWWW redirect redirects www requests to non www. For example, http://www.labstack.com will be redirected to http://labstack.com.
	// e.Pre(middleware.NonWWWRedirect())

	/*
		// Custom Configuration
		e.Use(middleware.HTTPSRedirectWithConfig(middleware.RedirectConfig{
		  Code: http.StatusTemporaryRedirect,
		}))

		//Configuration
		RedirectConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Status code to be used when redirecting the request.
		  // Optional. Default value http.StatusMovedPermanently.
		  Code int `json:"code"`
		}

		//Default Configuration
		DefaultRedirectConfig = RedirectConfig{
		  Skipper: DefaultSkipper,
		  Code:    http.StatusMovedPermanently,
		}
	*/

	// #Request ID Middleware
	// Request ID middleware generates a unique id for a request.

	e.Use(middleware.RequestID())

	/*
		// Custom Configuration
		e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		  Generator: func() string {
		    return customGenerator()
		  },
		}))

		//Configuration
		RequestIDConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Generator defines a function to generate an ID.
		  // Optional. Default value random.String(32).
		  Generator func() string
		}

		//Default Configuration
		DefaultRequestIDConfig = RequestIDConfig{
		  Skipper:   DefaultSkipper,
		  Generator: generator,
		}
	*/

	// #Rewrite Middleware
	// Rewrite middleware rewrites the URL path based on provided rules. It can be helpful for backward compatibility or just creating cleaner and more descriptive links.

	e.Pre(middleware.Rewrite(map[string]string{
		"/old":              "/new",
		"/api/*":            "/$1",
		"/js/*":             "/public/javascripts/$1",
		"/users/*/orders/*": "/user/$1/order/$2",
	}))

	/*
		// Custom Configuration
		e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{}))

		//Configuration
		// RewriteConfig defines the config for Rewrite middleware.
		RewriteConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Rules defines the URL path rewrite rules.
		  Rules map[string]string `yaml:"rules"`
		}

		//Default Configuration
		Name	Value
		Skipper	DefaultSkipper
	*/

	// #Secure Middleware
	// Secure middleware provides protection against cross-site scripting (XSS) attack, content type sniffing, clickjacking, insecure connection and other code injection attacks.

	e.Use(middleware.Secure())

	/*
		// Custom Configuration
		e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
			XSSProtection:         "",
			ContentTypeNosniff:    "",
			XFrameOptions:         "",
			HSTSMaxAge:            3600,
			ContentSecurityPolicy: "default-src 'self'",
		}))

		//Configuration
		SecureConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // XSSProtection provides protection against cross-site scripting attack (XSS)
		  // by setting the `X-XSS-Protection` header.
		  // Optional. Default value "1; mode=block".
		  XSSProtection string `yaml:"xss_protection"`

		  // ContentTypeNosniff provides protection against overriding Content-Type
		  // header by setting the `X-Content-Type-Options` header.
		  // Optional. Default value "nosniff".
		  ContentTypeNosniff string `yaml:"content_type_nosniff"`

		  // XFrameOptions can be used to indicate whether or not a browser should
		  // be allowed to render a page in a <frame>, <iframe> or <object> .
		  // Sites can use this to avoid clickjacking attacks, by ensuring that their
		  // content is not embedded into other sites.provides protection against
		  // clickjacking.
		  // Optional. Default value "SAMEORIGIN".
		  // Possible values:
		  // - "SAMEORIGIN" - The page can only be displayed in a frame on the same origin as the page itself.
		  // - "DENY" - The page cannot be displayed in a frame, regardless of the site attempting to do so.
		  // - "ALLOW-FROM uri" - The page can only be displayed in a frame on the specified origin.
		  XFrameOptions string `yaml:"x_frame_options"`

		  // HSTSMaxAge sets the `Strict-Transport-Security` header to indicate how
		  // long (in seconds) browsers should remember that this site is only to
		  // be accessed using HTTPS. This reduces your exposure to some SSL-stripping
		  // man-in-the-middle (MITM) attacks.
		  // Optional. Default value 0.
		  HSTSMaxAge int `yaml:"hsts_max_age"`

		  // HSTSExcludeSubdomains won't include subdomains tag in the `Strict Transport Security`
		  // header, excluding all subdomains from security policy. It has no effect
		  // unless HSTSMaxAge is set to a non-zero value.
		  // Optional. Default value false.
		  HSTSExcludeSubdomains bool `yaml:"hsts_exclude_subdomains"`

		  // ContentSecurityPolicy sets the `Content-Security-Policy` header providing
		  // security against cross-site scripting (XSS), clickjacking and other code
		  // injection attacks resulting from execution of malicious content in the
		  // trusted web page context.
		  // Optional. Default value "".
		  ContentSecurityPolicy string `yaml:"content_security_policy"`
		}

		//Default Configuration
		DefaultSecureConfig = SecureConfig{
		  Skipper:            DefaultSkipper,
		  XSSProtection:      "1; mode=block",
		  ContentTypeNosniff: "nosniff",
		  XFrameOptions:      "SAMEORIGIN",
		}
	*/

	// #Session Middleware
	// Session middleware facilitates HTTP session management backed by gorilla/sessions. The default implementation provides cookie and filesystem based session store; however, you can take advantage of community maintained implementation for various backends.

	/*
		Dependencies
		import (
		  "github.com/gorilla/sessions"
		  "github.com/labstack/echo-contrib/session"
		)

		e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

		e.GET("/", func(c echo.Context) error {
		  sess, _ := session.Get("session", c)
		  sess.Options = &sessions.Options{
		    Path:     "/",
		    MaxAge:   86400 * 7,
		    HttpOnly: true,
		  }
		  sess.Values["foo"] = "bar"
		  sess.Save(c.Request(), c.Response())
		  return c.NoContent(http.StatusOK)
		})

		// Custom Configuration
		e.Use(session.MiddlewareWithConfig(session.Config{}))

		//Configuration
		Config struct {
		  // Skipper defines a function to skip middleware.
		  Skipper middleware.Skipper

		  // Session store.
		  // Required.
		  Store sessions.Store
		}

		//Default Configuration
		DefaultConfig = Config{
		  Skipper: DefaultSkipper,
		}
	*/

	// #Static Middleware
	// Static middleware can be used to serve static files from the provided root directory.
	// This serves static files from static directory. For example, a request to /js/main.js will fetch and serve static/js/main.js file.

	e.Use(middleware.Static("/static"))

	/*
		// Custom Configuration
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		  Root:   "static",
		  Browse: true,
		}))

		//Configuration
		StaticConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Root directory from where the static content is served.
		  // Required.
		  Root string `json:"root"`

		  // Index file for serving a directory.
		  // Optional. Default value "index.html".
		  Index string `json:"index"`

		  // Enable HTML5 mode by forwarding all not-found requests to root so that
		  // SPA (single-page application) can handle the routing.
		  // Optional. Default value false.
		  HTML5 bool `json:"html5"`

		  // Enable directory browsing.
		  // Optional. Default value false.
		  Browse bool `json:"browse"`
		}

		//Default Configuration
		DefaultStaticConfig = StaticConfig{
		  Skipper: DefaultSkipper,
		  Index:   "index.html",
		}
	*/

	// #Trailing Slash Middleware

	// Add Trailing Slash
	// Add trailing slash middleware adds a trailing slash to the request URI.
	// e.Pre(middleware.AddTrailingSlash())

	// Remove Trailing Slash
	// Remove trailing slash middleware removes a trailing slash from the request URI.
	// e.Pre(middleware.RemoveTrailingSlash())

	/*
		// Custom Configuration
		e.Use(middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{
		  RedirectCode: http.StatusMovedPermanently,
		}))

		//Configuration
		TrailingSlashConfig struct {
		  // Skipper defines a function to skip middleware.
		  Skipper Skipper

		  // Status code to be used when redirecting the request.
		  // Optional, but when provided the request is redirected using this code.
		  RedirectCode int `json:"redirect_code"`
		}

		//Default Configuration
		DefaultTrailingSlashConfig = TrailingSlashConfig{
		  Skipper: DefaultSkipper,
		}
	*/

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

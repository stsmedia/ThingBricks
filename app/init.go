package app

import (
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/controllers"
	"github.com/stsmedia/thingbricks/app/persistence"
	"os"
	"strings"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		CorsFilter,
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		//		revel.SessionFilter,           // Restore and write the session cookie.
		//		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,  // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,        // Resolve the requested language
		HeaderFilter,            // Add some security based headers
		revel.InterceptorFilter, // Run interceptors around the action.
		revel.CompressFilter,    // Compress the result.
		revel.ActionInvoker,     // Invoke the action.
	}

	persistence.InitDb()

	var app controllers.AccountController

	if err := inject.Populate(&app); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var CorsFilter = func(c *revel.Controller, fc []revel.Filter) {
	if c.Request.Method == "OPTIONS" {
		c.Response.Out.Header().Add("Access-Control-Allow-Methods", strings.Join(c.Request.Header["Access-Control-Request-Method"], ","))
		c.Response.Out.Header().Add("Access-Control-Allow-Headers", strings.Join(c.Request.Header["Access-Control-Request-Headers"], ","))
		c.Response.Out.Header().Add("Access-Control-Allow-Origin", strings.Join(c.Request.Header["Origin"], ","))
		c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")
		c.RenderJson("{foo}")
	} else {
		fc[0](c, fc[1:])
	}
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {

	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Access-Control-Allow-Origin", strings.Join(c.Request.Header["Origin"], ","))
	c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

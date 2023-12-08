package app

import (
	"expense-track/app/controllers"
	"expense-track/app/db"
	_ "github.com/lib/pq"
	_ "github.com/revel/modules"
	"github.com/revel/revel"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(db.InitDB)
	// revel.OnAppStart(FillCache)

	revel.FilterController(controllers.App{}).Insert(AuthFilter, revel.BEFORE, revel.ActionInvoker)
	revel.FilterController(controllers.Auth{}).Insert(NoAuthFilter, revel.BEFORE, revel.ActionInvoker)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

var AuthFilter = func(c *revel.Controller, fc []revel.Filter) {
	if _, ok := c.Session["username"]; !ok {
		c.Result = c.Redirect("/auth")
		return
	}
	fc[0](c, fc[1:])
}
var NoAuthFilter = func(c *revel.Controller, fc []revel.Filter) {
	if _, ok := c.Session["username"]; ok {
		c.Result = c.Redirect("/")
		return
	}
	fc[0](c, fc[1:])
}

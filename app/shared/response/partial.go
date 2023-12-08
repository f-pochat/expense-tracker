package response

import "github.com/revel/revel"

func RenderPartial(c *revel.Controller, partial string, args map[string]interface{}) revel.Result {
	for key, value := range args {
		c.ViewArgs[key] = value
	}
	return c.RenderTemplate(partial + ".html")
}

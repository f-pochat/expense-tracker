package controllers

import (
	"expense-track/app/repositories"
	"expense-track/app/shared/response"
	"github.com/revel/revel"
	"strings"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) GetCategories() revel.Result {
	userID, _ := c.Session["id"]
	categories, err := repositories.GetCategories(userID.(string))
	if err != nil {
		c.Log.Error(err.Error())
		return c.RenderTemplate("sidebar/SidebarItems.html")
	}
	var lowercaseCategories []map[string]interface{}
	for _, category := range categories {
		lowercaseCategories = append(lowercaseCategories, map[string]interface{}{
			"Name": category,
			"Href": strings.ToLower(category),
		})
	}

	args := map[string]interface{}{
		"categories": lowercaseCategories,
	}
	return response.RenderPartial(c.Controller, "sidebar/SidebarItems", args)
}

package controllers

import (
	"expense-track/app/models/dto"
	"expense-track/app/repositories"
	"expense-track/app/shared/request"
	"expense-track/app/shared/response"
	"expense-track/app/shared/utils"
	"github.com/revel/revel"
)

type Auth struct {
	*revel.Controller
}

func (c Auth) Index() revel.Result {
	return c.Render()
}

func (c Auth) Register() revel.Result {
	return c.Render()
}

func (c Auth) Login() revel.Result {
	login := dto.LoginDTO{}
	if err := request.BindForm(c.Params.Form, &login); err != nil {
		c.Log.Error(err.Error())
		args := map[string]interface{}{
			"error": "Invalid User",
		}
		return response.RenderPartial(c.Controller, "Auth/LoginForm", args)
	}
	user, err := repositories.GetUser(login.Username)
	if err != nil {
		c.Log.Error(err.Error())
		args := map[string]interface{}{
			"error": "User not found",
		}
		return response.RenderPartial(c.Controller, "Auth/LoginForm", args)
	}
	if user.Password != utils.EncodePassword(login.Password) {
		c.Log.Error(err.Error())
		args := map[string]interface{}{
			"error": "Invalid password",
		}
		return response.RenderPartial(c.Controller, "Auth/LoginForm", args)
	}
	// Redirects on success
	c.Response.Out.Header().Add("HX-Location", "/")
	c.Session["username"] = user.Username
	c.Session["id"] = user.ID
	return c.RenderJSON(user)
}

func (c Auth) RegisterUser() revel.Result {
	register := dto.RegisterDTO{}
	if err := request.BindForm(c.Params.Form, &register); err != nil {
		c.Log.Error(err.Error())
		args := map[string]interface{}{
			"error": "Invalid User",
		}
		return response.RenderPartial(c.Controller, "Auth/RegisterForm", args)
	}
	register.Password = utils.EncodePassword(register.Password)
	newUser, err := repositories.CreateUser(register)
	if err != nil {
		c.Log.Error(err.Error())
		args := map[string]interface{}{
			"error": "User or email already exists",
		}
		return response.RenderPartial(c.Controller, "Auth/RegisterForm", args)
	}

	// Redirects on success
	c.Response.Out.Header().Add("HX-Location", "/")
	c.Session["username"] = newUser.Username
	c.Session["id"] = newUser.ID
	return c.RenderJSON(newUser)
}

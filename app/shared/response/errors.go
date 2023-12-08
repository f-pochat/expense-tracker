package response

import (
	"github.com/revel/revel"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func RenderJSONError(c *revel.Controller, status int, errMessage string) revel.Result {
	c.Response.Status = status
	c.Response.ContentType = "application/json"

	errorJSON := ErrorResponse{
		Error:   http.StatusText(status),
		Message: errMessage,
	}

	return c.RenderJSON(errorJSON)
}

package middlewares

import (
	"net/http"
	"strings"

	"github.com/alanpv92/events/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {

	authorization := ctx.Request.Header.Get("Authorization")
	if authorization == "" {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.ErrorResponse("no token present"),
		)
	}
	tokenString := strings.Split(authorization, " ")[1]
	data, err := helpers.VerifyJwtToken(tokenString)
	if err != nil && data != nil && ctx.FullPath() == "/auth/refresh" {
		ctx.Set("id", data["id"])
		ctx.Set("email", data["email"])
		ctx.Next()
		return
	}
	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			helpers.ErrorResponse("invaild token"),
		)
		return
	}
	ctx.Set("id", data["id"])
	ctx.Set("email", data["email"])
	ctx.Next()
}

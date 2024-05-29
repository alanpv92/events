package controller

import (
	"net/http"

	"github.com/alanpv92/events/database"
	"github.com/alanpv92/events/helpers"
	"github.com/alanpv92/events/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.InvaildBodyErrorResponse())
		return
	}
	err = user.ValidateLoginBody()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
		return
	}

	dbUser, err := database.GetUserByEmail(user.Email)
	if err != nil {

		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	if dbUser == nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("user has not registred"))
		return
	}
	isPasswordOk := helpers.VerifyPassword(user.Password, dbUser.Password)
	if !isPasswordOk {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invaild password"))
		return
	}
	token, err := helpers.GenerateToken(*dbUser, false)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	refreshToken, err := helpers.GenerateToken(*dbUser, true)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	dbUser.RefreshToken = refreshToken
	dbUser.Token = token
	ctx.JSON(http.StatusCreated, gin.H{
		"data": dbUser.AuthResponse(),
	})
}

func Register(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.InvaildBodyErrorResponse())
		return
	}
	err = user.ValidateRegisterBody()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(err.Error()))
	}

	dbUser, err := database.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	if dbUser != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("user already is already registred"))
		return
	}

	hashedPassword, err := helpers.HashPasswod(user.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	user.Password = hashedPassword
	id, err := database.InsertUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	user.Id = id
	token, err := helpers.GenerateToken(user, false)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	refreshToken, err := helpers.GenerateToken(user, true)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}

	user.Token = token
	user.RefreshToken = refreshToken
	ctx.JSON(http.StatusCreated, gin.H{
		"data": user.AuthResponse(),
	})

}

func RefreshToken(ctx *gin.Context) {
	id, isIdPresent := ctx.Get("id")
	email, isEmailPresent := ctx.Get("email")
	if !isIdPresent || !isEmailPresent {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invaild claims"))
		return
	}
	var user models.User
	user.Id = id.(string)
	user.Email = email.(string)
	err := ctx.ShouldBindBodyWithJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invaild refresh token"))
		return
	}
	err = user.ValidateRefreshToken()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse(" refresh token is required"))
		return
	}
	_, err = helpers.VerifyJwtToken(user.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invaild refresh token"))
		return
	}
	token, tokenError := helpers.GenerateToken(user, false)
	refreshToken, refreshTokenError := helpers.GenerateToken(user, true)
	if tokenError != nil || refreshTokenError != nil {
		ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("could not generate tokens"))
		return
	}
	user.Token = token
	user.RefreshToken = refreshToken
	ctx.JSON(http.StatusAccepted, user.TokenResponse())

}

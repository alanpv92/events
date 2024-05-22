package controller

import (
	"net/http"

	"github.com/alanpv92/events/database"
	"github.com/alanpv92/events/helpers"
	"github.com/alanpv92/events/models"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	//check request body
	//check if user exists
	//check if password matches
	//generate jwt token
	//send
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
 	isPasswordOk:= helpers.VerifyPassword(user.Password, dbUser.Password)
	if !isPasswordOk{
		ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invaild password"))
		return;
	}
	token,err:=helpers.GenerateToken(*dbUser);
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}
	dbUser.Token=token;
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
	token, err := helpers.GenerateToken(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helpers.RandomErrorResponse())
		return
	}

	user.Token = token
	ctx.JSON(http.StatusCreated, gin.H{
		"data": user.AuthResponse(),
	})

	//check if user has already registred
	//hash the password
	//insert user into database
	//generate jwt
	//send sucess response

}

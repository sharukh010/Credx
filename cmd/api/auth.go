package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sharukh010/credx/internal/auth"
	"github.com/sharukh010/credx/internal/store"
)

type Gender int 
const (
	Male Gender = iota
	Female
)
type userRegisterRequest struct {
	UserName string `json:"user_name" binding:"required,min=5"`
	Name name `json:"name"`
	Gender Gender `json:"gender" binding:"omitempty,min=0,max=1"`
	Email string `json:"email" binding:"required,email"`
	DOB string `json:"dob" binding:"required,min=10,max=10"`// dd/mm/yyyy format 
	Password string `json:"password" binding:"required,min=4"` // change it while deploying min=8
}

type userLoginRequest struct {
	UserName string `json:"user_name" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=4"` // change it while deploying min=8
}

type userLoginResponse struct {
	 Token string `json:"token"`
}

type name struct {
	FirstName string `json:"first_name" binding:"required,min=5"`
	LastName string `json:"last_name" binding:"required,min=5"`
}

// userRegistrationHandler godoc
// @Summary Register user
// @Description Registers a new user.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body userRegisterRequest true "Registration payload"
// @Success 201 {object} UserDataResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/register [post]
func (app *application) userRegistrationHandler(c *gin.Context){
	r := userRegisterRequest{}

	if err := c.BindJSON(&r); err != nil {
		badRequestResponse(c,err)
		return 
	}

	user := &store.User{
		UserName: r.UserName,
		Name: store.Name{
			FirstName: r.Name.FirstName,
			LastName: r.Name.LastName,
		},
		Email: r.Email,
		DOB: r.DOB,
	}
	
	if r.Gender == 0 {
		user.Gender = "male"
	}else{
		user.Gender = "female"
	}

	hash,err := store.HashPassword(r.Password)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}
	user.Password = hash 
	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	if err := app.store.Users.Create(ctx,user); err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	jsonResponse(c,http.StatusCreated,user)
}

// userLoginHandler godoc
// @Summary Login user
// @Description Authenticates a user and returns a JWT token.
// @Tags auth
// @Accept json
// @Produce json
// @Param request body userLoginRequest true "Login payload"
// @Success 200 {object} LoginDataResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/log-in [get]
func (app *application) userLoginHandler(c *gin.Context){
	r := userLoginRequest{}

	if err := c.BindJSON(&r); err != nil {
		badRequestResponse(c,err)
		return 
	}

	ctx,cancel := context.WithTimeout(c,DatabaseOperationsTimeOut)
	defer cancel()

	user,err := app.store.Users.GetByUserName(ctx,r.UserName)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			notFoundResponse(c,err)
			return 
		default:
			internalServerErrorResponse(c,err)			
		}
		return 
	}

	if err := store.CompareHashAndPassword(user.Password,r.Password); err != nil {
		invalidCredentialsResponse(c,err)
		return 
	}

	claims := auth.Claims{
		UserID : user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),//change this while deploying 
			Issuer: "credx",
		},

	}
	token,err := auth.GenerateJWT(claims,app.config.JWTSecret)
	if err != nil {
		internalServerErrorResponse(c,err)
		return 
	}

	response := userLoginResponse{
		Token: token,
	}

	jsonResponse(c,http.StatusOK,response)
}

func (app *application) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			unauthorizedResponse(c,fmt.Errorf("%s Missing",USERTOKEN))
			c.Abort()
			return 
		}

		tokenStr := strings.TrimPrefix(authHeader,"Bearer ")

		claims,err := auth.ParseJWT(tokenStr,app.config.JWTSecret)
		if err != nil {
			unauthorizedResponse(c,fmt.Errorf("%s Invalid",USERTOKEN))
			c.Abort()
			return 
		}

		c.Set("userID",claims.UserID)
		c.Set("email",claims.Email)
		c.Next()
	}
}

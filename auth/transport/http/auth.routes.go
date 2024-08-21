package http

import (
	//"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"spectator.main/domain"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(r *gin.RouterGroup, uu domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: uu,
	}
	r.POST("/auth/signup", handler.SignUp)
	r.POST("/auth/login", handler.Login)
}

func (ah *AuthHandler) Login(ctx *gin.Context) {
	var request domain.LoginRequest
	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	user, err := ah.AuthUsecase.GetUserByEmail(ctx, request.Email)

	if err != nil {
		ctx.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	secret_key := os.Getenv("SECRET_KEY")
	expiry_time := os.Getenv("EXPIRY_TIME")
	expiry, _ := strconv.Atoi(expiry_time)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	// 	return
	// }

	accessToken, err := ah.AuthUsecase.CreateAccessToken(user, secret_key, expiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := ah.AuthUsecase.CreateRefreshToken(user, secret_key, expiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

func (ah *AuthHandler) SignUp(ctx *gin.Context) {
	var request domain.SignupRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = ah.AuthUsecase.GetUserByEmail(ctx, request.Email)

	if err == nil {

		ctx.JSON(http.StatusConflict, domain.ErrorResponse{Message: "user already exists "})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = ah.AuthUsecase.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	secret_key := os.Getenv("SECRET_KEY")
	expiry_time := os.Getenv("EXPIRY_TIME")
	expiry, _ := strconv.Atoi(expiry_time)
	accessToken, err := ah.AuthUsecase.CreateAccessToken(&user, secret_key, expiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := ah.AuthUsecase.CreateRefreshToken(&user, secret_key, expiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, signupResponse)
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/recor-glitch/zepo-backend/internal/auth"
	"github.com/recor-glitch/zepo-backend/internal/domain/user"
	"github.com/recor-glitch/zepo-backend/internal/infrastructure/db"
	"github.com/recor-glitch/zepo-backend/internal/infrastructure/repository"
	usecase_user "github.com/recor-glitch/zepo-backend/internal/usecase/user"
	"github.com/recor-glitch/zepo-backend/internal/utils"
)

var userRepo = repository.NewUserRespository(db.ConnectDB())

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	u, err := usecase_user.GetByID(id, userRepo)
	if err != nil {
		utils.MapDBError(err, c)
		return
	}

	c.JSON(http.StatusOK, u)
}

func CreateUser(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindBodyWithJSON((&u)); err != nil {
		utils.MapDBError(err, c)
		return
	}

	er := usecase_user.CreateUser(&u, userRepo)
	if er != nil {
		utils.MapDBError(er, c)
		return
	}

	access_token, acc_err := auth.GenerateAccessToken(u.Id)
	if acc_err != nil {
		utils.MapDBError(acc_err, c)
		return
	}

	refresh_token, ref_err := auth.GenerateRefreshToken(u.Id)
	if ref_err != nil {
		utils.MapDBError(ref_err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"access_token": access_token, "refresh_token": refresh_token})
}

func InvalidateAccessToken(c *gin.Context) {
	var tokens user.Tokens
	if err := c.ShouldBindBodyWithJSON(&tokens); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error(), "statuscode": http.StatusBadRequest})
		return
	}

	claims, err := auth.ValidateRefreshToken(tokens.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error(), "statuscode": http.StatusUnauthorized})
		return
	}

	access_token, acc_err := auth.GenerateAccessToken(claims.ID)
	if acc_err != nil {
		utils.MapDBError(acc_err, c)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"access_token": access_token, "refresh_token": tokens.AccessToken})
}

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusCreated, gin.H{})
}

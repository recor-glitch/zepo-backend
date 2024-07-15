package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/recor-glitch/zepo-backend/internal/domain"
	"gorm.io/gorm"
)

func MapDBError(err error, c *gin.Context) {
	if err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": domain.ErrInternal.Error(), "statusCode": http.StatusInternalServerError})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"msg": domain.ErrRecordNotFound.Error(), "statusCode": http.StatusNotFound})
		return
	}

	if pgErr, ok := err.(*pgconn.PgError); ok {
		switch pgErr.Code {
		// duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)
		case "23505":
			c.JSON(http.StatusBadGateway, gin.H{"msg": domain.ErrDuplicateEntry.Error(), "statusCode": http.StatusBadGateway})
		default:
			c.JSON(http.StatusBadGateway, gin.H{"msg": domain.ErrInternal.Error(), "statusCode": http.StatusBadGateway})
		}
	} else {
		// HANDLE NON-SQL ERRORS
		c.JSON(http.StatusInternalServerError, gin.H{"msg": domain.ErrInternal.Error(), "statusCode": http.StatusInternalServerError})
	}

}

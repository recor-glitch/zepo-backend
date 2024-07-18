package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
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

	if pqErr, ok := err.(*pq.Error); ok {
		fmt.Printf("%+v", pqErr)

		switch pqErr.Code {
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

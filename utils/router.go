package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/snykk/go_gin_ca/constants"
	"github.com/snykk/go_gin_ca/datatransfers"
)

func AuthOnly(c *gin.Context) {
	if !c.GetBool(constants.IsAuthenticatedKey) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, datatransfers.Response{Message: "user not authenticated"})
	}
}

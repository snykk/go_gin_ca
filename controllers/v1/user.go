package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/snykk/go_gin_ca/constants"
	"github.com/snykk/go_gin_ca/datatransfers"
	"github.com/snykk/go_gin_ca/handlers"
	"github.com/snykk/go_gin_ca/models"
)

func GETUser(c *gin.Context) {
	var err error
	var userInfo datatransfers.UserInfo
	if err = c.ShouldBindUri(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	var user models.User
	if user, err = handlers.Handler.RetrieveUser(userInfo.Id); err != nil {
		c.JSON(http.StatusNotFound, datatransfers.Response{Status: false, Message: "cannot find user"})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: "user data fetched successfully", Data: datatransfers.UserInfo{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}})
}

func PUTUser(c *gin.Context) {
	var err error
	var user datatransfers.UserUpdate

	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	if err = handlers.Handler.UpdateUser(c.GetUint(constants.UserIDKey), user); err != nil {
		c.JSON(http.StatusNotModified, datatransfers.Response{Status: false, Message: "failed updating user"})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: "user data updated successfully", Data: user})
}

package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/snykk/go_gin_ca/datatransfers"
	"github.com/snykk/go_gin_ca/handlers"
)

func POSTLogin(c *gin.Context) {
	var err error
	var user datatransfers.UserLogin

	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	var token string
	if token, err = handlers.Handler.AuthenticateUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Status: false, Message: "incorrect username or password"})
		return
	}

	c.JSON(http.StatusOK, datatransfers.Response{Status: true, Message: "login success", Data: fmt.Sprintf("Token: %s", token)})
}

func POSTRegister(c *gin.Context) {
	var err error
	var user datatransfers.UserSignup

	if err = c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, datatransfers.Response{Status: false, Message: err.Error()})
		return
	}

	if err = handlers.Handler.RegisterUser(user); err != nil {
		c.JSON(http.StatusUnauthorized, datatransfers.Response{Status: false, Message: "failed registering user"})
		return
	}

	c.JSON(http.StatusCreated, datatransfers.Response{Status: true, Message: "user created successfuly", Data: user})
}

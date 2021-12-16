package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/xapponi/golang-microservices/mvc/services"
	"github.com/xapponi/golang-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			Status:  http.StatusBadRequest,
			Code:    "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}
	user, apiErr := services.UsersService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
	}
	utils.Respond(c, http.StatusOK, user)
}


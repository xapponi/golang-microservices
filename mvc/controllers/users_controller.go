package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/xapponi/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := (strconv.ParseInt(userIdParam, 10, 64))
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number"))
		// Return the Bad request to the client
		return
	}
	user, err := services.GetUser(userId)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// Handle the err and return to the client
		return
	}
	// return user to client
	jsonValue,_ := json.Marshal(user)
	resp.Write(jsonValue)

}
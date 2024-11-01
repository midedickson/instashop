package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/midedickson/instashop/dto"
	"github.com/midedickson/instashop/utils"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Create user logic
	var createUserPayload dto.UserAuthPayload
	err := json.NewDecoder(r.Body).Decode(&createUserPayload)
	if err != nil {
		log.Printf("Error decoding create user payload: %v", err)
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	if !createUserPayload.Validate() {
		utils.Dispatch400Error(w, "Invalid Payload, email and password is required", nil)
		return
	}
	user, err := c.userUseCase.CreateUser(createUserPayload)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	utils.Dispatch200(w, "User created successfully", user)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	// Login user logic
	var loginPayload dto.UserAuthPayload
	err := json.NewDecoder(r.Body).Decode(&loginPayload)
	if err != nil {
		log.Printf("Error decoding login payload: %v", err)
		utils.Dispatch400Error(w, "Invalid Payload", err)
		return
	}
	if !loginPayload.Validate() {
		utils.Dispatch400Error(w, "Invalid Payload, email and password is required", nil)
		return
	}

	// userClaim := c.userUseCase
}

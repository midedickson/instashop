package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/midedickson/instashop/internal/dto"
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
	user, err := c.userService.CreateUser(createUserPayload)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	log.Println(user, err, "create user result")
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

	userClaim, err := c.userService.GetUserByEmail(loginPayload.Email)
	if err != nil {
		utils.Dispatch404Error(w, "User not found", nil)
		return
	}
	if !c.userService.VerifyUserPasswordWithHash(loginPayload.Password) {
		utils.Dispatch403Error(w, "Invalid password", nil)
		return
	}

	if !userClaim.IsActive {
		utils.Dispatch403Error(w, "User is not active", nil)
		return
	}

	accessToken, err := c.userService.GenerateJwtTokenForUser(userClaim)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}

	userDetails := &dto.LoginResponse{
		UserDetails: userClaim,
		AccessToken: accessToken,
	}
	utils.Dispatch200(w, "Logged in successfully", userDetails)
}

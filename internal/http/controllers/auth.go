package controllers

import (
	"log"
	"net/http"

	"github.com/midedickson/instashop/constants"
	"github.com/midedickson/instashop/internal/dto"
	"github.com/midedickson/instashop/utils"
)

func (c *Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Create user logic
	createUserPayload := r.Context().Value(constants.SignupPayloadCtxKey{}).(dto.UserAuthPayload)

	user, err := c.userService.CreateUser(createUserPayload)
	if err != nil {
		utils.Dispatch500Error(w, err)
		return
	}
	log.Println(user, err, "create user result")
	utils.Dispatch200(w, "User created successfully", user)
}

func (c *Controller) ActivateUser(w http.ResponseWriter, r *http.Request) {
	// Activate user logic
	activateUserPayload := r.Context().Value(constants.ActivateUserPayloadCtxKey{}).(dto.ActivateUserPayload)
	user, err := c.userService.ActivateUser(activateUserPayload)
	if err != nil {
		utils.Dispatch404Error(w, "User not found or activation code is invalid", nil)
		return
	}
	utils.Dispatch200(w, "User activated successfully", user)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	// Login user logic

	loginPayload := r.Context().Value(constants.LoginPayloadCtxKey{}).(dto.UserAuthPayload)

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

func (c *Controller) AdminLogin(w http.ResponseWriter, r *http.Request) {
	// Login user logic

	loginPayload := r.Context().Value(constants.LoginPayloadCtxKey{}).(dto.UserAuthPayload)

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

	if userClaim.Role == constants.ADMIN_ROLE {
		utils.Dispatch403Error(w, "You are not authorized to access this endpoint", nil)
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

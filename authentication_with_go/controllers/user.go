package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		userService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("id")
	fmt.Println("user id we got: ", userId)
	if userId == "" {
		userId = r.Context().Value("userId").(string)
	}

	if userId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user id in query parameters"))
		return
	}

	user, err := uc.userService.GetUserByID(userId)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Failed to get user", fmt.Errorf("user with id %s not found", userId))
		return
	}

	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User retrieved successfully", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(*dto.CreateUserRequestDTO)

	user, err := uc.userService.CreateUser(payload)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "User creation failed", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusCreated, "User created successfully", user)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("payload").(*dto.LoginUserRequestDTO)

	token, err := uc.userService.LoginUser(payload)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusUnauthorized, "Login failed", err)
		return
	}
	
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User logged in successfully", token)
}
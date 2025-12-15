package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	RoleId          uint
	PhoneNumber     string `json:"phoneNumber" validate:"required"`
	Email           string `json:"email" validate:"required,email "`
}

type UpdateRequest struct {
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password,omitempty"`
	ConfirmPassword string `json:"confirmPassword,omitempty"`
	RoleId          uint
	PhoneNumber     string `json:"phoneNumber" validate:"required"`
	Email           string `json:"email" validate:"required,email "`
}

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Username    string    `json:"username"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phoneNumber"`
	Email       string    `json:"email"`
}

type LoginResponse struct {
	UserResponse UserResponse `json:"user"`
	Token        string       `json:"token"`
}

type RegisterResponse struct {
	UserResponse UserResponse `json:"user"`
}

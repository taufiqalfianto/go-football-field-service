package routes

import (
	"user-service/controllers"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IUserRoute interface {
	Run()
}

type NewUserRoute (controllers.IControllerRegistry)


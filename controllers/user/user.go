package user

import (
	"net/http"
	errWrap "user-service/common/error"
	"user-service/common/response"
	errConstant "user-service/constants/error"
	"user-service/domain/dto"
	"user-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	service services.IServiceRegistry
}

type IUserController interface {
	Login(*gin.Context)
	Register(*gin.Context)
	Update(*gin.Context)
	GetUserLogin(*gin.Context)
	GetUserByUUID(*gin.Context)
}

func NewUserController(service services.IServiceRegistry) IUserController {
	return &UserController{service: service}
}

func (u *UserController) Login(ctx *gin.Context) {
	request := &dto.LoginRequest{}

	err := ctx.ShouldBindJSON(request)

	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	user, err := u.service.GetUser().Login(ctx, request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusInternalServerError,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code:  http.StatusOK,
		Data:  user.UserResponse,
		Token: &user.Token,
		Gin:   ctx,
	})
}

func (u *UserController) Register(ctx *gin.Context) {
	request := &dto.RegisterRequest{}

	err := ctx.ShouldBindJSON(request)

	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	user, err := u.service.GetUser().Register(ctx, request)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == errConstant.ErrUsernameExist || err == errConstant.ErrEmailExist {
			statusCode = http.StatusConflict
		} else if err == errConstant.ErrPasswordDoesNotMatch {
			statusCode = http.StatusBadRequest
		}

		response.HttpResponse(response.ParamHTTPResp{
			Code: statusCode,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: user.UserResponse,
		Gin:  ctx,
	})
}

// GetUserByUUID implements [IUserController].
func (u *UserController) GetUserByUUID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	user, err := u.service.GetUser().GetUserByUUID(ctx, uuid)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == errConstant.ErrUserNotFound {
			statusCode = http.StatusNotFound
		}
		response.HttpResponse(response.ParamHTTPResp{
			Code: statusCode,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}

// GetUserLogin implements [IUserController].
func (u *UserController) GetUserLogin(ctx *gin.Context) {
	user, err := u.service.GetUser().GetUserLogin(ctx)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == errConstant.ErrUserNotFound {
			statusCode = http.StatusNotFound
		}
		response.HttpResponse(response.ParamHTTPResp{
			Code: statusCode,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}

// Update implements [IUserController].
func (u *UserController) Update(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	request := &dto.UpdateRequest{}

	err := ctx.ShouldBindJSON(request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResp{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errWrap.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusUnprocessableEntity,
			Message: &errMessage,
			Data:    errResponse,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	user, err := u.service.GetUser().Update(ctx, request, uuid)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err == errConstant.ErrUserNotFound {
			statusCode = http.StatusNotFound
		} else if err == errConstant.ErrUsernameExist || err == errConstant.ErrEmailExist {
			statusCode = http.StatusConflict
		} else if err == errConstant.ErrPasswordDoesNotMatch {
			statusCode = http.StatusBadRequest
		}

		response.HttpResponse(response.ParamHTTPResp{
			Code: statusCode,
			Err:  err,
			Gin:  ctx,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: user,
		Gin:  ctx,
	})
}

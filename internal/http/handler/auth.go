package handler

import (
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/service"
	"DynamicLED/internal/http/dto"
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (controller *Controller) Register(c *gin.Context) {
	var cred dto.Credentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if err := cred.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	err := controller.Service.Auth.Register(c.Request.Context(), cred.Login, cred.Password)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExist) {
			c.AbortWithStatusJSON(http.StatusConflict, dto.NewErrorResponse(err.Error()))
			return
		}

		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (controller *Controller) Login(c *gin.Context) {
	var cred dto.Credentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if err := cred.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	access, refresh, err := controller.Service.Auth.Login(c.Request.Context(), cred.Login, cred.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) ||
			errors.Is(err, service.ErrTokenExpired) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse(err.Error()))
			return
		}

		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.NewTokenPair(access, refresh))
}

func (controller *Controller) Refresh(c *gin.Context) {
	var tokens dto.TokenPair
	if err := c.ShouldBindJSON(&tokens); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	access, err := controller.Service.Auth.Refresh(c.Request.Context(), tokens.Access, tokens.Refresh)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) ||
			errors.Is(err, service.ErrTokenExpired) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.NewErrorResponse(err.Error()))
			return
		}

		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.NewTokenPair(access, tokens.Refresh))
}

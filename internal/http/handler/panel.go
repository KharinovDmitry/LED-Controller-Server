package handler

import (
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/service"
	"DynamicLED/internal/http/dto"
	"DynamicLED/internal/http/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func (controller *Controller) RegisterPanel(c *gin.Context) {
	_, _, userUUID, err := util.GetClaimsFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	var req dto.RegisterPanelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if err := controller.Service.RegisterPanel(c.Request.Context(), req.Rev, req.Mac, req.Host, userUUID); err != nil {
		if errors.Is(err, service.ErrPanelNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, dto.NewErrorResponse(err.Error()))
			return
		}
		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (controller *Controller) SendTaskToPanel(c *gin.Context) {
	panelUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	var req dto.PanelTask
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if err := controller.Service.SendTaskToPanel(c.Request.Context(), panelUUID, req.ToEntity()); err != nil {
		if errors.Is(err, service.ErrPanelNotRegistered) ||
			errors.Is(err, service.ErrPanelNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, dto.NewErrorResponse(err.Error()))
			return
		}

		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (controller *Controller) GetPanelByMAC(c *gin.Context) {
	mac := c.Param("mac")
	if mac == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse("mac is required"))
		return
	}

	panel, err := controller.Service.GetPanelByMac(c.Request.Context(), mac)
	if err != nil {
		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.Panel(panel))
}

func (controller *Controller) GetPanelByUUID(c *gin.Context) {
	panelUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	panel, err := controller.Service.GetPanelByUUID(c.Request.Context(), panelUUID)
	if err != nil {
		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.Panel(panel))
}

func (controller *Controller) GetPanelByUserUUID(c *gin.Context) {
	_, _, userUUID, err := util.GetClaimsFromContext(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	panels, err := controller.Service.GetPanelsByUserUUID(c.Request.Context(), userUUID)
	if err != nil {
		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.PanelsToDTO(panels))
}

func (controller *Controller) GetDisplay(c *gin.Context) {
	panelUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	display, err := controller.Service.GetPanelDisplayByUUID(c.Request.Context(), panelUUID)
	if err != nil {
		if errors.Is(err, service.ErrPanelNotRegistered) ||
			errors.Is(err, service.ErrPanelNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, dto.NewErrorResponse(err.Error()))
			return
		}

		if errors.Is(err, service.ErrCacheUpdate) {
			slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
			c.Status(http.StatusOK)
			return
		}

		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.DisplayToDTO(display))
}

func (controller *Controller) SyncDisplay(c *gin.Context) {
	panelUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	display, err := controller.Service.SyncPanelDisplay(c.Request.Context(), panelUUID)
	if err != nil {
		if errors.Is(err, service.ErrPanelNotRegistered) ||
			errors.Is(err, service.ErrPanelNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, dto.NewErrorResponse(err.Error()))
			return
		}

		slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, dto.DisplayToDTO(display))
}

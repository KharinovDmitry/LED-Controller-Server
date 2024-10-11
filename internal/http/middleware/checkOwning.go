package middleware

import (
	"DynamicLED/internal/domain/constant"
	"DynamicLED/internal/domain/repository"
	"DynamicLED/internal/domain/service"
	"DynamicLED/internal/http/dto"
	"DynamicLED/internal/http/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

func CheckPanelOwningByUUID(panelRepo repository.Panel) func(c *gin.Context) {
	return func(c *gin.Context) {
		panelUUID, err := uuid.Parse(c.Param("uuid"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
			return
		}

		panel, err := panelRepo.GetPanelByUUID(c.Request.Context(), panelUUID)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, service.ErrPanelNotFound)
				return
			}

			slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		_, _, userUUID, err := util.GetClaimsFromContext(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
			return
		}

		if panel.Owner != userUUID {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

func CheckPanelOwningByMac(panelRepo repository.Panel) func(c *gin.Context) {
	return func(c *gin.Context) {
		panelMac := c.Param("mac")
		if panelMac == "" {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		panel, err := panelRepo.GetPanelByMac(c.Request.Context(), panelMac)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, service.ErrPanelNotFound)
				return
			}

			slog.Error(c.FullPath(), slog.String(constant.ErrorField, err.Error()))
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		_, _, userUUID, err := util.GetClaimsFromContext(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
			return
		}

		if panel.Owner != userUUID {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Next()
	}
}

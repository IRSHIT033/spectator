package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"spectator.main/domain"
	"spectator.main/internals/bootstrap"
)

type ConfigHandler struct {
	ConfigUsecase domain.ConfigUsecase
	config        *bootstrap.Config
}

func NewConfigHandler(cfg *bootstrap.Config, r *gin.RouterGroup, uu domain.ConfigUsecase) {
	handler := &ConfigHandler{
		ConfigUsecase: uu,
		config:        cfg,
	}
	r.POST("/config", handler.CreateConfig)
	r.GET("/config/:user_id", handler.GetConfigByUserID)
	r.PUT("/config/:config_id/site", handler.AddSiteConfig)
	r.DELETE("/config/:config_id/site", handler.RemoveSiteConfig)
	r.PATCH("/config/:config_id/site", handler.UpdateSiteConfig)
}

func (h *ConfigHandler) CreateConfig(c *gin.Context) {
	var config domain.ConfigDetails
	if err := c.ShouldBindJSON(&config); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ConfigUsecase.InsertOne(c, &config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func (h *ConfigHandler) GetConfigByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	config, err := h.ConfigUsecase.GetByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, config)
}

func (h *ConfigHandler) AddSiteConfig(c *gin.Context) {
	var siteConfig domain.SiteConfig
	if err := c.ShouldBindJSON(&siteConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ConfigUsecase.AddSiteConfig(c, &siteConfig, c.Param("config_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Site config added successfully"})
}

func (h *ConfigHandler) RemoveSiteConfig(c *gin.Context) {
	var removeConfig domain.RemoveConfigRequest
	if err := c.ShouldBindJSON(&removeConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ConfigUsecase.RemoveSiteConfig(c, removeConfig.SiteUrl, c.Param("config_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Site config removed successfully"})
}

func (h *ConfigHandler) UpdateSiteConfig(c *gin.Context) {
	var siteConfig domain.SiteConfig
	if err := c.ShouldBindJSON(&siteConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.ConfigUsecase.UpdateSiteConfig(c, &siteConfig, c.Param("config_id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Site config updated successfully"})
}

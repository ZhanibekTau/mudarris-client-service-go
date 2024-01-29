package handlers

import (
	"client-service-go/app/entityManager"
	structures "client-service-go/config/configStruct"
	ginhelper "github.com/exgamer/go-rest-sdk/pkg/helpers/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Handler struct {
	manager *entityManager.Manager
	appData *structures.AppData
}

func NewHandler(manager *entityManager.Manager, appData *structures.AppData) *Handler {
	return &Handler{manager: manager, appData: appData}
}

// @title Subbox API
// @version 1.0

func (h *Handler) Init() *gin.Engine {
	router := ginhelper.InitRouter(h.appData.AppConfig)

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	healthRouter := router.Group("/health")
	{
		healthRouter.GET("alive", healthcheck)
		healthRouter.GET("ready", healthcheck)
	}

	hook := router.Group("client-service-go")
	{
		v1 := hook.Group("v1")
		{
			endpoint := v1.Group("client")
			{
				endpoint.POST("create", h.create)
			}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "true", "status": http.StatusOK})
}

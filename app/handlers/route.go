package handlers

import (
	ginhelper "github.com/exgamer/go-rest-sdk/pkg/helpers/gin"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"user-service-go/app/entityManager"
	structures "user-service-go/config/configStruct"
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

	hook := router.Group("user-service-go")
	{
		v1 := hook.Group("v1")
		{
			authClient := v1.Group("client")
			{
				endpoint := authClient.Group("auth")
				{
					endpoint.POST("create", h.create)
					endpoint.POST("login", h.login)
				}
				api := authClient.Group("api", h.clientIdentity)
				{
					api.PUT("update", h.update)
					api.GET("get-by-id", h.getById)
					api.GET("get-ustaz-by-id", h.getByIdUstaz)
				}
			}

			authUstaz := v1.Group("ustaz")
			{
				endpointUstaz := authUstaz.Group("auth")
				{
					endpointUstaz.POST("create", h.createUstaz)
					endpointUstaz.POST("login", h.loginUstaz)
				}
				apiUstaz := authUstaz.Group("api", h.ustazIdentity)
				{
					apiUstaz.PUT("update", h.updateUstaz)
					apiUstaz.GET("get-by-id", h.getByIdUstaz)
				}
			}

			//course := v1.Group("course", h.clientIdentity)
			//{
			//	course.GET("get-all", h.getAll)
			//	course.POST("enrollment", h.enrollment)
			//}
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "true", "status": http.StatusOK})
}

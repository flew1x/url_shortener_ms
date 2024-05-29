package httpv1

import (
	"log/slog"
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/flew1x/url_shortener_auth_ms/internal/cache"
	"github.com/flew1x/url_shortener_auth_ms/internal/config"
	"github.com/flew1x/url_shortener_auth_ms/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger  *slog.Logger
	service *service.Service
	config  *config.Config
	cache   *cache.Cache
}

func NewHandler(logger *slog.Logger, service *service.Service, config *config.Config, cache *cache.Cache) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
		config:  config,
		cache:   cache,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use()

	floatRate := float64(h.config.ServerConfig.GetLimitPerSecond())
	limiter := tollbooth.NewLimiter(floatRate, &limiter.ExpirableOptions{
		DefaultExpirationTTL: time.Second,
	})

	common := router.Group("/", tollbooth_gin.LimitHandler(limiter))
	{
		common.GET("/:url", h.redirectToOriginalURL)

		api := common.Group("/api")
		{
			v1 := api.Group("/v1")
			{
				v1.GET("/healthcheck", h.healthcheck)
				v1.POST("/shorten", h.shortenURL)
			}
		}
	}

	return router
}

package httpv1

import "github.com/gin-gonic/gin"

func (h *Handler) healthcheck(c *gin.Context) {
	c.Status(200)
}

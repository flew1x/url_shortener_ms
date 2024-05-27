package http_v1

import (
	"net/http"

	"github.com/flew1x/url_shortener_auth_ms/internal/service"
	"github.com/gin-gonic/gin"
)

type GetShortenURLParams struct {
	URL string `json:"url"`
}

type GetShortenUrlResponse struct {
	ShortURL string `json:"short_url"`
}

// shortenURL is the HTTP handler for the "/shorten-url" endpoint.
// It receives a JSON object containing the URL to shorten.
// It returns a JSON object containing the shortened URL.
//
// Parameters:
// - c: the gin.Context for the operation.
func (h *Handler) shortenURL(c *gin.Context) {
	var request GetShortenURLParams
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithError(http.StatusBadRequest, ErrInvalidRequest)
		return
	}

	if request.URL == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrRequiredUrl)
		return
	}

	shortURL, err := h.service.UrlShortener.Create(c.Request.Context(), request.URL)
	if err != nil {
		if err == service.ErrNotValidURL {
			c.AbortWithStatusJSON(http.StatusBadRequest, service.ErrNotValidURL)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, ErrInternalError)
		return
	}

	c.JSON(http.StatusOK, GetShortenUrlResponse{ShortURL: shortURL})
}

// redirectToOriginalURL is the HTTP handler for the "/{shorten_url}" endpoint.
// It redirects the client to the original URL associated with the given short URL.
//
// Parameters:
// - c: the gin.Context for the operation.
func (h *Handler) redirectToOriginalURL(c *gin.Context) {
	shortURL := c.Param(SHORTEN_URL_PARAM)
	if shortURL == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrRequiredUrl)
		return
	}

	originalURL, err := h.service.UrlShortener.GetByShort(c.Request.Context(), shortURL)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrInternalError)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, originalURL.Origin)
}

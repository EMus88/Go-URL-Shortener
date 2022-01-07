package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}
type ShortURL struct {
	URL string `json:"url"`
}
type Result struct {
	Result string `json:"result"`
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) HandlerGet(c *gin.Context) {
	id := c.Param("id")
	longURL := h.service.GetURL(id)
	if longURL == "" {
		c.String(http.StatusBadRequest, "URL not found")
		return
	}
	c.Status(http.StatusTemporaryRedirect)
	c.Header("Location", longURL)
}

func (h *Handler) HandlerPostText(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil || len(body) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	id := h.service.SaveURL(string(body))
	baseURL := os.Getenv("BASE_URL")
	c.String(http.StatusCreated, baseURL+id)

}
func (h *Handler) HandlerPostJSON(c *gin.Context) {
	if c.GetHeader("content-type") != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not allowed content-type"})
		return
	}
	var ShortURL ShortURL
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(jsonData, &ShortURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := h.service.SaveURL(ShortURL.URL)
	baseURL := os.Getenv("BASE_URL")
	longURL := baseURL + id
	var result Result
	result.Result = longURL
	c.JSON(http.StatusCreated, result)
}

package handler

import (
	"io/ioutil"
	"net/http"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}
type URL struct {
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
	if c.GetHeader("content-type") != "text/plain" {
		c.String(http.StatusBadRequest, "Not allowed content-type")
		return
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil || len(body) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	id := h.service.SaveURL(string(body))
	c.String(http.StatusCreated, "http://localhost:8080/"+id)

}
func (h *Handler) HandlerPostJSON(c *gin.Context) {
	if c.GetHeader("content-type") != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not allowed content-type"})
		return
	}
	var Url URL
	if err := c.ShouldBindJSON(&Url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := h.service.SaveURL(Url.URL)
	longURL := "http://localhost:8080/" + id
	var result Result
	result.Result = longURL
	c.JSON(http.StatusCreated, result)
}

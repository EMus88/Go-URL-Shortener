package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/app/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}
type LongURL struct {
	URL string `json:"url"`
}
type Result struct {
	Result string `json:"result"`
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

//=================================================================
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

//==================================================================
func (h *Handler) HandlerPostText(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil || len(body) == 0 {
		c.Status(http.StatusBadRequest)
		return
	}
	id := h.service.SaveURL(string(body))
	c.String(http.StatusCreated, h.service.Config.BaseURL+"/"+id)
}

//===================================================================
func (h *Handler) HandlerPostJSON(c *gin.Context) {
	if c.GetHeader("content-type") != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not allowed content-type"})
		return
	}
	var longURL LongURL
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(jsonData, &longURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := h.service.SaveURL(longURL.URL)
	shortURL := h.service.Config.BaseURL + "/" + id
	var result Result
	result.Result = shortURL
	c.JSON(http.StatusCreated, result)
}

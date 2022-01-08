package handler

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/EMus88/go-musthave-shortener-tpl/internal/app/service"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository"
	"github.com/EMus88/go-musthave-shortener-tpl/internal/repository/models/file"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/magiconair/properties/assert"
)

func TestHandler_HandlerPostText(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name        string
		requestBody string
		want        want
	}{
		{
			name:        "test 1",
			requestBody: "https://yandex.ru/search/?text=go&lr=11351&clid=9403sdfasdfasdfasdf",
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name:        "test 2",
			requestBody: "",
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	//load env variables
	if err := godotenv.Load("/omhe/emus/Рабочий стол/Yandex/go-musthave-shortener-tpl/.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var model file.Model
			r := repository.NewStorage()
			s := service.NewService(r, &model)
			h := NewHandler(s)

			gin.SetMode(gin.ReleaseMode)
			router := gin.Default()

			router.POST("/", h.HandlerPostText)

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.requestBody))
			w := httptest.NewRecorder()
			req.Header.Set("content-type", "text/plain")
			router.ServeHTTP(w, req)
			result := w.Result()
			defer result.Body.Close()
			assert.Equal(t, result.StatusCode, tt.want.statusCode)

		})
	}
}

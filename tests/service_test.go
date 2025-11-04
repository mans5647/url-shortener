package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"url-shortener/handlers"
	"url-shortener/models/url"
	"github.com/labstack/echo/v4"
)

type Service struct {
	handler *handlers.ShortenerHandler
	*echo.Echo
}

func NewService() *Service   {

	handler := handlers.NewMemorySequenceShortener()
	
	e := echo.New()
	e.POST("/api/v1/shorten", handler.HandleShorten)

	return &Service{
		handler: handlers.NewMemorySequenceShortener(),
		Echo: e,
	}
}

func helperToJsonAsReader(v any) io.Reader {
	buf, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return bytes.NewReader(buf)
}

func TestBadRequest(t * testing.T) {
	
	service := NewService()

	recorder := httptest.NewRecorder()

	uri := url.PlainUrl{Url: ""}

	req := httptest.NewRequest("POST", "/api/v1/shorten", helperToJsonAsReader(&uri))
	req.Header.Add("Content-Type", "application/json")

	service.ServeHTTP(recorder, req)
	t.Run("Empty url, returns 400", func(t *testing.T) {

		if recorder.Code != http.StatusBadRequest {
			t.Error("expected 400, but got:", recorder.Code)
		}
	})

	service.ServeHTTP(recorder, httptest.NewRequest("POST", "/api/v1/shorten", strings.NewReader("Hello World!")))

	t.Run("Incorrect body, returns 400", func(t *testing.T) {
		if recorder.Code != http.StatusBadRequest {
			t.Error("expected 400, but got:", recorder.Code)
		}
	})

	

}
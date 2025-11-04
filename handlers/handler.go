package handlers

import (
	"errors"
	"net/http"
	"strings"
	"url-shortener/database"
	"url-shortener/models/url"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


type ShortenerHandler struct {
	storage database.Storage
}

func NewShortenerHandler(storage database.Storage) * ShortenerHandler {
	return &ShortenerHandler{
		storage: storage,
	}
}

func NewMemorySequenceShortener() * ShortenerHandler {
	return &ShortenerHandler{
		storage: database.NewMemoryStorage(),
	}
}

func NewDatabaseSequenceShortener(dbName, collName, host string, port int) (* ShortenerHandler, error) {

	store, err := database.NewMongoStorage(dbName, collName, host, port)

	if err != nil {
		return nil, err
	}

	return &ShortenerHandler{
		storage: store,
	}, nil
}

func (handler * ShortenerHandler) HandleShorten(c echo.Context) error {

	oldUrl := url.PlainUrl{}

	if !strings.HasPrefix(c.Request().Header.Get("Content-Type"), "application/json") {
		return c.String(http.StatusBadRequest, "invalid content type")
	}

	if err := c.Bind(&oldUrl); err != nil {
		return c.String(http.StatusBadRequest, "body error")
	}
	
	if len(oldUrl.Url) == 0 {
		return c.String(http.StatusBadRequest, "empty url")
	}


	nextId := handler.storage.GenerateId(c.Request().Context())
	newUrl, err := url.Shorten(c.Request().Context(), nextId, url.DefaultCutSize, &oldUrl)

	if err != nil {
		return c.String(http.StatusInternalServerError, "failed to generate url")
	}

	errInsert := handler.storage.Store(c.Request().Context(), newUrl)

	if errInsert != nil {
		c.JSON(http.StatusInternalServerError, struct{
			description string
			errorCode int
		}{
			description: "insert error",
			errorCode: -1,
		})
	}

	c.Response().Header().Set("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
	c.Response().Header().Set("Access-Control-Allow-Methods", "POST")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	return c.JSON(http.StatusOK, newUrl)
}


func (handler * ShortenerHandler) HandleRedirect(c echo.Context) error {

	code := c.Param("code")

	uri, err := handler.storage.FindByCode(c.Request().Context(), code)

	if errors.Is(err, mongo.ErrNoDocuments) {
		return c.String(http.StatusNotFound, "not found")
	}

	return c.Redirect(http.StatusMovedPermanently, uri.Url)

}
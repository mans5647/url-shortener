package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"url-shortener/database"
	"url-shortener/models"
)

const (
	jsonContentType = "application/json"	
	htmlContentType = "text/html"
)


func toJson(obj any) (string, error) {
	bytes, err := json.Marshal(obj)

	if err != nil {
		return "", errors.New("failed to convert to json")
	}

	return string(bytes), nil
}

func fromJson(bytes []byte, obj any) error {
	return json.Unmarshal(bytes, obj)
} 

func doResponse(status int, body string, contentType string, w http.ResponseWriter) {
	
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	w.Write([]byte(body))
}

func responseAsJson(status int, body string, w http.ResponseWriter) {

	doResponse(status, body, jsonContentType, w)
}

func responseAsHTML(status int, body string, w http.ResponseWriter) {

	doResponse(status, body, htmlContentType, w)
}

/*
	shortens url using 
*/
func ShortenUrlHandler(w http.ResponseWriter, r * http.Request) {
	
	var requestBody []byte
	var urlToConvert models.OldFormUrl

	MimeChecker := func (val string) bool {
		return strings.HasPrefix(val, jsonContentType)
	}


	if r.Method != http.MethodPost {
		responseAsHTML(http.StatusMethodNotAllowed, "<h1>Method is not allowed!</h1>", w)
		return
	}
	
	// check if body is json type

	if value := r.Header.Get("Content-Type"); !MimeChecker(value) {
		responseAsHTML(http.StatusBadRequest, "<h1>Invalid request</h1>", w)
		return
	}

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		responseAsHTML(http.StatusBadRequest, "<h1>Read body failed</h1>", w)
		return
	}

	

	if fromJson(requestBody, &urlToConvert) != nil {
		responseAsHTML(http.StatusBadRequest, "<h1>Failed to process request body</h1>", w)
		return
	}

	shortenedUrl, err := database.AddNewShortUrl(&urlToConvert)
	
	if err != nil {
		responseAsHTML(http.StatusInternalServerError, "<h1>Failed to short url!</h1>", w)
		return
	}

	json, _ := toJson(shortenedUrl)
	responseAsJson(http.StatusOK, json, w)
	r.Body.Close()
}

/*
	Handles short url redirection by using "Location" header in response

	for example:
	http://localhost:8080/1Hg redirects 
	to https://some.real.long.url//making-things-right

*/
func RedirectByLinkCodeHandler(w http.ResponseWriter, r * http.Request) {

	code := r.PathValue("code")
	realUrl, err := database.FindRealUrlByCode(code)
	if err != nil {
		responseAsHTML(http.StatusInternalServerError, "could't find short url", w)
		return
	}

	w.Header().Add("Location", realUrl)
	doResponse(http.StatusFound, "", htmlContentType, w)
}

/*
	removes all url records from table in database
*/
func DeleteAllUrlsHandler(w http.ResponseWriter, r * http.Request) {

	if r.Method != http.MethodDelete {
		responseAsHTML(http.StatusMethodNotAllowed, "", w)
		return
	}

	if database.DeleteAllUrls() == nil {
		responseAsHTML(http.StatusNoContent, "", w)
		return
	}


	responseAsHTML(http.StatusInternalServerError, "", w)

}
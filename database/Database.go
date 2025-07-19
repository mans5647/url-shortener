package database

import (
	"errors"
	"url-shortener/models"
	"url-shortener/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn * gorm.DB = nil

var DefaultDsn = "host=localhost user=mans password=psql_123 dbname=urls_db port=5432 sslmode=disable"

func OpenPostgresConnection(dsn string) bool {

	newConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return false
	}

	/* assign pointer to connection */
	conn = newConn
	return true
}

func AutoMigrateTables() bool {

	if conn == nil {
		return false
	}

	err := conn.Table("urls").AutoMigrate(&models.NewFormUrl{})
	return err == nil
}

func AddNewShortUrl(urlForm * models.OldFormUrl) (*models.NewFormUrl, error) {


	maxId := 0
	shortUrl := &models.NewFormUrl{}
	// 1. get max id

	
	conn.Table("urls").Select("max(urls.id)").Scan(&maxId)
	
	code := utils.GenerateShortCode(maxId, urlForm.Url)

	shortUrl.Code = code
	shortUrl.RealUrl = urlForm.Url

	// 3. add record to database

	if conn.Table("urls").Save(shortUrl).Error != nil {
		return nil, errors.New("failed to save value")
	}

	return shortUrl, nil
}

func FindRealUrlByCode(code string) (string, error) {

	var realUrl string
	if conn == nil {
		return "", errors.New("no connection")
	}

	err := conn.Table("urls").
		Select("real_url").
		Where("code = ?", code).
		Limit(1).
		Scan(&realUrl).Error

	if err != nil {
		return "", errors.New("error in finding url")
	}

	return realUrl, nil
}

func DeleteAllUrls() (err error) {

	if conn == nil {
		return errors.New("no connection")
	}

	conn.Exec("DELETE from urls")
	return
}
package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"path"

	"github.com/Mohammed-Aadil/ds-storage/config"
	"github.com/jinzhu/gorm"
)

//SaveFile save file to destination
func SaveFile(file multipart.File, header *multipart.FileHeader) error {
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(config.StaticPath, header.Filename), data, 0666); err != nil {
		return err
	}
	return nil
}

//GetDBConnection return a DB connection
func GetDBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName),
	)
	// db.DB().SetMaxIdleConns(5)
	// db.DB().SetMaxOpenConns(10)
	// db.DB().SetConnMaxLifetime(time.Minute * 5)
	if err != nil {
		return nil, err
	}
	return db, nil
}

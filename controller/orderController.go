package controller

import (
	"challenge-godb/config"
	"challenge-godb/model"
	"time"
)

func CreateOrder(order model.Order) error {
	db := config.ConnectDb()
	defer db.Close()

	return nil
}

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

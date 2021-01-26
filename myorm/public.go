package myorm

import (
	"log"
	"testing"

	"gorm.io/gorm"
)

func close(db *gorm.DB) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recoverd in ", r)
		}
	}()
	if db != nil {
		conn, err := db.DB()
		if err != nil {
			log.Println("close error ", err)
			return
		}
		if err := conn.Close(); err != nil {
			log.Println("close error ", err)
			return
		}
	}
}

func handleRecover(t *testing.T) {
	if r := recover(); r != nil {
		t.Error(r)
	}
}

package main

import (
	"errors"

	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/external"
	"github.com/ITEBARPLKelompok3/peminjaman-ruangan/backend/model"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func main() {
	// Prepare Migration
	db, err := external.GetPostgresClient()
	if err != nil {
		sentry.CaptureException(err)
	}

	tx := db.Begin()
	tx.AutoMigrate(getModels()...)
	defer tx.Rollback()

	if err := tx.First(&model.Role{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		// Insert seed data
		tx.Create([]*model.Role{
			{
				RoleName: "admin",
			},
			{
				RoleName: "user",
			},
		})
	}
	err = tx.Commit().Error
	if err != nil {
		log.Error(err)
		sentry.CaptureException(err)
		return
	}
	log.Info("Schema migrated!")
	sentry.CaptureMessage("Schema migrated!")
}

func getModels() (models []interface{}) {
	models = []interface{}{
		model.User{},
		model.Role{},
	}
	return
}

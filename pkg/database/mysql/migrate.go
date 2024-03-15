package mysql

import (
	"intern-bcc/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.Role{},
		&entity.User{},
		&entity.Post{},
		&entity.Uni{},
		&entity.District{},
	); err != nil {
		return err
	}
	return nil
}

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
		&entity.Minat{},
		&entity.Skill{},
	); err != nil {
		return err
	}
	return nil
}

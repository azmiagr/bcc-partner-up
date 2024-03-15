package entity

type Uni struct {
	ID   uint
	Name string `json:"name" gorm:"type:varchar(255);not null;"`
	User []User
}

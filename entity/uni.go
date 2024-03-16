package entity

type Uni struct {
	ID    uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Name  string `json:"name" gorm:"type:varchar(255);not null;"`
	Users []User
}

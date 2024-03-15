package entity

type Skill struct {
	ID   uint
	Name string
	User []*User `json:"user" gorm:"many2many:user_skill;"`
}

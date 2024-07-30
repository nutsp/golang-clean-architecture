package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:pass"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) IsNil() bool {
	return u == nil
}

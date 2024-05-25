package models

type UserRole struct {
	UserID uint `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	RoleID uint `gorm:"primaryKey;autoIncrement:false" json:"role_id"`
	User   User `gorm:"foreignKey:UserID"`
	Role   Role `gorm:"foreignKey:RoleID"`
}

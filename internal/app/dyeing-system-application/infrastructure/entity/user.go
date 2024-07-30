package entity

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment;comment:主键" json:"id"`
	Name      string    `gorm:"size:20;not null;comment:用户名" json:"name"`
	Phone     string    `gorm:"size:15;not null;comment:手机号" json:"phone"`
	Email     string    `gorm:"size:30;not null;comment:邮箱地址" json:"email"`
	Address   string    `gorm:"size:100;not null;comment:联系住址" json:"address"`
	Password  string    `gorm:"size:20;not null;comment:密码" json:"password"`
	CreatedAt time.Time //`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time //`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time //`json:"deleted_at,omitempty"`
	Status    int16     `gorm:"not null;comment:状态 1-正常，2-注销" json:"status"`
	Bak       string    `gorm:"size:100;comment:备注" json:"bak"`
}

package entity

import "time"

type Dye struct {
	ID          uint64    `gorm:"primary_key;auto_increment;comment:主键" json:"id"`
	Name        string    `gorm:"size:20;not null;comment:染料名" json:"name"`
	TotalAmount float32   `gorm:"default:0;comment:总量" json:"total_amount"`
	LastAmount  float32   `gorm:"default:0;comment:上次总量" json:"last_amount"`
	Phone       string    `gorm:"size:15;default:0;comment:染料公司手机号" json:"phone"`
	Company     string    `gorm:"size:100;default: ;comment:染料公司" json:"company"`
	Address     string    `gorm:"size:100;default: ;comment:染料公司地址" json:"address"`
	Status      int16     `gorm:"default:1;comment:状态 1-正常，2-注销" json:"status"`
	Bak         string    `gorm:"size:100;comment:备注" json:"bak"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   time.Time `gorm:"default:NULL" json:"deleted_at"`
}

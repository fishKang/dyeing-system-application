package entity

import "time"

type Channel struct {
	ID        uint64    `gorm:"primary_key;auto_increment;comment:主键" json:"id"`
	SerialNum string    `gorm:"unique;size:40;not null;comment:请求编号" json:"serialnum"`
	Zoneno    string    `gorm:"size:10;not null;comment:地区号" json:"zoneno"`
	User      string    `gorm:"size:50;comment:用户名" json:"user"`
	Service   string    `gorm:"size:50;not null;comment:服务名" json:"service"`
	Method    string    `gorm:"size:50;not null;comment:方法名" json:"method"`
	Request   string    `gorm:"size:1000;not null;comment:入参" json:"request"`
	Response  string    `gorm:"size:2000;not null;comment:出参" json:"response"`
	Workdate  string    `gorm:"size:10;not null;comment:调用日期" json:"workdate"`
	Worktime  string    `gorm:"size:8;not null;comment:调用时间" json:"worktime"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:NULL" json:"deleted_at"`
}

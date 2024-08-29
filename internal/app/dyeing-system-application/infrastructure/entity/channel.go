package entity

import "time"

type Channel struct {
	ID        uint64    `gorm:"primary_key;auto_increment;comment:主键" json:"id"`
	SerialNum string    `gorm:"unique;size:40;not null;comment:请求编号" json:"serialnum"`
	Zoneno    string    `gorm:"size:10;not null;comment:地区号" json:"zoneno"`
	User      string    `gorm:"size:50;comment:用户名" json:"user"`
	Service   string    `gorm:"size:50;not null;comment:服务名" json:"service"`
	Method    string    `gorm:"size:50;not null;comment:方法名" json:"method"`
	Request   string    `gorm:"size:200;not null;comment:入参" json:"request"`
	Response  string    `gorm:"size:500;not null;comment:出参" json:"response"`
	Workdate  string    `gorm:"size:10;not null;comment:调用日期" json:"workdate"`
	Worktime  string    `gorm:"size:8;not null;comment:调用时间" json:"worktime"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:NULL" json:"deleted_at"`
}

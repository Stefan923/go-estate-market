package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id uint `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone; not null;"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone; not null;"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone; not null;"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UTC()
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return
}

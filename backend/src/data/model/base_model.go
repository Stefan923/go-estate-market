package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id uint `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone; not null;"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone; null;"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone; null;"`
}

func (m *BaseModel) BeforeCreate(*gorm.DB) error {
	m.CreatedAt = time.Now().UTC()
	return nil
}

func (m *BaseModel) BeforeUpdate(*gorm.DB) error {
	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return nil
}

func (m *BaseModel) BeforeDelete(*gorm.DB) error {
	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return nil
}

package model

import (
	"database/sql"
	"time"
)

type BaseModel struct {
	Id uint `gorm:"primarykey"`

	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone; not null;"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone; not null;"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone; not null;"`
}

func (m *BaseModel) BeforeCreate() {
	m.CreatedAt = time.Now().UTC()
	return
}

func (m *BaseModel) BeforeUpdate() {
	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return
}

func (m *BaseModel) BeforeDelete() {
	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	return
}

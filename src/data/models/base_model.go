package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	CreatedAt  time.Time    `gorm:"type:TIMESTAMP with time zone;not null"`
	ModifiedAt sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime `gorm:"type:TIMESTAMP with time zone;null"`

	CreatedBy      *sql.NullInt64 `gorm:"null"`
	LastModifiedBy *sql.NullInt64 `gorm:"null"`
	DeletedBy      *sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}

	if value != nil {
		UserId = &sql.NullInt64{Int64: value.(int64), Valid: true}
	}

	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = UserId
	return
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}

	if value != nil {
		UserId = &sql.NullInt64{Int64: value.(int64), Valid: true}
	}

	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.LastModifiedBy = UserId
	return
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	value := tx.Statement.Context.Value("UserId")
	var UserId = &sql.NullInt64{Valid: false}

	if value != nil {
		UserId = &sql.NullInt64{Int64: value.(int64), Valid: true}
	}

	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = UserId
	return
}

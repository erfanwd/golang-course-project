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
	var userId = &sql.NullInt64{Valid: false}

	if value != nil {
		switch v := value.(type) {
		case int64:
			userId = &sql.NullInt64{Int64: v, Valid: true}
		case float64:
			userId = &sql.NullInt64{Int64: int64(v), Valid: true}
		case int:
			userId = &sql.NullInt64{Int64: int64(v), Valid: true}
		}
	}

	m.CreatedAt = time.Now().UTC()
	m.CreatedBy = userId
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

// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID                  string               `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	CreatedAt           time.Time            `gorm:"column:createdAt;type:timestamp without time zone;not null;default:now()" json:"createdAt"`
	DeletedAt           *time.Time           `gorm:"column:deletedAt;type:timestamp without time zone" json:"deletedAt"`
	UserAuthentications []UserAuthentication `json:"user_authentications"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

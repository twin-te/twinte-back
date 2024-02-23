// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTag = "tags"

// Tag mapped from table <tags>
type Tag struct {
	ID       string `gorm:"column:id;type:uuid;primaryKey" json:"id"`
	UserID   string `gorm:"column:user_id;type:uuid;not null;uniqueIndex:UQ_b861cf6ec9af09190780481a311,priority:1" json:"user_id"`
	Name     string `gorm:"column:name;type:text;not null" json:"name"`
	Position int32  `gorm:"column:position;type:integer;not null;uniqueIndex:UQ_b861cf6ec9af09190780481a311,priority:2" json:"position"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
// Code generated by codegen/idtype/generate.py; DO NOT EDIT.

package idtype

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type UserID uuid.UUID

func (id UserID) String() string {
	return uuid.UUID(id).String()
}

func (id UserID) IsZero() bool {
	return uuid.UUID(id) == uuid.Nil
}

func (id UserID) Less(other UserID) bool {
	for i := 0; i < 16; i++ {
		if id[i] == other[i] {
			continue
		}
		return id[i] < other[i]
	}
	return false
}

func (id *UserID) StringPtr() *string {
	if id == nil {
		return nil
	}
	return lo.ToPtr(id.String())
}

func NewUserID() UserID {
	return UserID(uuid.New())
}

func ParseUserID(s string) (UserID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return UserID{}, fmt.Errorf("failed to parse UserID %#v", s)
	}
	return UserID(id), nil
}
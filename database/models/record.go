package models

import (
	"math/big"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Description string
	Type RecordType
	Amount big.Float
	Reference time.Time
	Paid bool `gorm:"default:false"`
}

type RecordType uint

const (
	Revenue RecordType = iota
	Debt
)

func (recordType RecordType) String() string {
	return [...]string{"revenue", "debt"}[recordType]
}
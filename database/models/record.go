package models

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Description string     `json:"description"`
	Type        RecordType `json:"type"`
	Amount      float32    `json:"amount" gorm:"type:decimal(10,4);"`
	Reference   time.Time  `json:"reference"`
	Paid        bool       `json:"paid" gorm:"default:false"`
}

type RecordType uint

const (
	Revenue RecordType = iota
	Debt
)

func recordTypeMapping() map[string]int {
	recordTypeMap := make(map[string]int)
	recordTypeMap["revenue"] = int(Revenue)
	recordTypeMap["debt"] = int(Debt)

	return recordTypeMap
}

func ParseRecordType(recordType string) (RecordType, error) {
	recordTypeMap := recordTypeMapping()

	normalizedValue := strings.TrimSpace(strings.ToLower(recordType))
	mappedValue, ok := recordTypeMap[normalizedValue]

	if !ok {
		return RecordType(0), errors.New("invalid type")
	}

	return RecordType(mappedValue), nil
}

func (recordType *RecordType) String() string {
	recordTypeMap := recordTypeMapping()

	var label string

	for i, rt := range recordTypeMap {
		if rt == int(*recordType) {
			label = i
			break
		}
	}

	return label
}

func (recordType *RecordType) MarshalJSON() ([]byte, error) {
	return json.Marshal(recordType.String())
}

func (recordType *RecordType) UnmarshalJSON(data []byte) error {
	var value string

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	record, err := ParseRecordType(value)

	if err != nil {
		return err
	}

	*recordType = record
	return nil
}

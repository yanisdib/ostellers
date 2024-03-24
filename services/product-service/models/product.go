package models

import (
	"strings"
	"time"
)

type Availability string

const (
	IN_STOCK         Availability = "IN STOCK"
	OUT_OF_STOCK     Availability = "OUT OF STOCK"
	IN_PROVISION     Availability = "IN PROVISION"
	NOT_YET_RELEASED Availability = "NOT YET RELEASED"
)

type ProductFormat string

const (
	PHYSICAL ProductFormat = "PHYSICAL"
	DIGITAL  ProductFormat = "DIGITAL"
	UNKNOW   ProductFormat = "UNKNOWN"
)

// TODO: add value objects
type Product struct {
	ID           string
	Label        string
	Description  string
	Stock        uint16
	Availability Availability
	Formats      []ProductFormat
	ReleasedAt   time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Categories   map[int]*Category
}

func StringToProductFormat(s string) ProductFormat {
	s = strings.ToUpper(s)
	switch s {
	case string(PHYSICAL):
		return PHYSICAL
	case string(DIGITAL):
		return DIGITAL
	default:
		return UNKNOW
	}
}

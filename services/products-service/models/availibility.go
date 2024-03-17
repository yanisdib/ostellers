package models

type Availability string

const (
	IN_STOCK         Availability = "IN_STOCK"
	OUT_OF_STOCK     Availability = "OUT_OF_STOCK"
	IN_PROVISION     Availability = "IN_PROVISION"
	NOT_YET_RELEASED Availability = "NOT_YET_RELEASED"
)

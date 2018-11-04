package models

// Order status
const (
	APIPathOrder = "/orders"

	OrderStatusUnassigned = "UNASSIGNED"
	OrderStatusTaken      = "TAKEN"

	ErrorOrderNotFound         = "ORDER_NOT_FOUND"
	ErrorOrderAlreadyBeenTaken = "ORDER_ALREADY_BEEN_TAKEN"
	ErrorDescription           = "ERROR_DESCRIPTION"
	ErrorInvalidParameters     = "INVALID_PARAMETERS"
	ErrorCalculateFailed       = "CALCULATE_FAILED"

	SQLTimeFormat = "2006-01-02 15:04:05"
)

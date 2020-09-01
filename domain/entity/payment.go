package entity

import "time"

// Payment structure required for payment calculation
type Payment struct {
	BorrowerPaymentAmount         float64
	Date                          time.Time
	InitialOutstandingPrincipal   float64
	Interest                      float64
	Principal                     float64
	RemainingOutstandingPrincipal float64
}

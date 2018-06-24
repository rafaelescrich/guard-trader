package operation

import (
	"time"
)

// Operation info
type Operation struct {
	orderID        int64
	symbol         string
	status         int
	amount         float64
	startTimestamp time.Time
	qt             int
	price          string
	time           string
}

func (op *Operation) GetSymbol() string {
	return op.symbol
}

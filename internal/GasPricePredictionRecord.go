package internal

import "time"

type GasPricePredictionRecord struct {
	Id                   string
	RecordType           int
	GasPrice             uint64
	GasBase              uint64
	GasUsed              uint64
	MaxPriority          uint64
	OriginalBlockHeight  uint64
	PredictedBlockHeight uint64
	CreateTime           time.Time
	UpdateTime           time.Time
}

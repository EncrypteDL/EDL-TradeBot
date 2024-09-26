package db

// Reciept Stores a receipt's data.
type Reciept struct {
	Symbol    string
	TxType    string
	Qauntity  float64
	Price     float32
	Profit    float32
	TimeStamp int64
}

// DB is a Database interface
type DB interface {
	GetLastReceiptes(limit int) ([]*Reciept, error)
	InsertReceipt(receipt *Reciept) error
	Close() error
}


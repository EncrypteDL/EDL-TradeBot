package etcd

import (
	"encoding/json"
	"time"

	"github.com/EncrypteDL/EDL-TradeBot/internal/db"
	"go.etcd.io/bbolt"
)

type DB struct {
	conn *bbolt.DB
}

var tradingPairSymbol string

// Open opens a new BoltDB connection with the provided URL and creates a bucket with the symbol name if it doesn't exist.
func Open(url, symbol string) (database *DB, err error) {
	conn, err := bbolt.Open(url, 0600, &bbolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	// Create a bucket if it doesn't exist
	err = conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(symbol))
		return err
	})

	if err != nil {
		return nil, err
	}

	tradingPairSymbol = symbol
	return &DB{conn: conn}, nil
}

// Close closes the BoltDB connection
func (d *DB) Close() error {
	return d.conn.Close()
}

// GetLastReceiptes retrieves the last 'limit' number of receipts from the BoltDB, ordered from most recent to oldest
func (d *DB) GetLastReceiptes(limit int) ([]*db.Reciept, error) {
	var receipts []*db.Reciept

	err := d.conn.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(tradingPairSymbol))
		if b == nil {
			return nil // if bucket doesn't exist, return empty result
		}

		c := b.Cursor()

		// Iterate from the last entry to the first
		for k, v := c.Last(); k != nil && len(receipts) < limit; k, v = c.Prev() {
			var receipt db.Reciept
			if err := json.Unmarshal(v, &receipt); err != nil {
				return err
			}
			receipts = append(receipts, &receipt)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return receipts, nil
}

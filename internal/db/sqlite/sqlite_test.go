package sqlite

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/EncrypteDL/EDL-TradeBot/internal/db"
)

var Dbb *DB

var receipts = []*db.Reciept{
	{
		Symbol:    "BTCUSDT",
		TxType:    "BUY",
		Qauntity:  1.0,
		Profit:    1.0,
		Price:     1.0,
		TimeStamp: 1,
	},
	{
		Symbol:    "BTCUSDT",
		TxType:    "SELL",
		Qauntity:  1.0,
		Profit:    1.0,
		Price:     1.0,
		TimeStamp: 2,
	},
}

func TestMain(m *testing.M) {
	var err error
	Dbb, err = Open(":memory:", "BTCUSDT", context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer Dbb.Close()

	os.Exit(m.Run())
}

func TestInsertReceipt(t *testing.T) {
	err := Dbb.InsertReceipt(receipts[1])
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLastReceipts(t *testing.T) {
	err := Dbb.InsertReceipt(receipts[0])
	if err != nil {
		t.Fatal(err)
	}

	n := 2

	r, err := Dbb.GetLastReceipts(n)
	if err != nil {
		t.Fatal(err)
	}

	if len(receipts) == 0 {
		t.Fatalf("no receipts")
	}

	if len(receipts) != n {
		t.Fatalf("got %d, want %d", len(receipts), n)
	}

	for i, receipt := range receipts {
		if receipt == r[i] {
			t.Errorf("got %v, want %v", r[i], receipt)
		}
	}
}

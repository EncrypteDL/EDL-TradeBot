package etcd

// import (
// 	"encoding/json"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/EncrypteDL/EDL-TradeBot/internal/db"
// 	"go.etcd.io/bbolt"
// )

// func TestOpenAndClose(t *testing.T) {
// 	// Create a temporary database file
// 	testDB := "test_open_close.db"
// 	defer os.Remove(testDB) // clean up after the test

// 	// Test opening the database
// 	dbInstance, err := Open(testDB, "BTC-USD")
// 	if err != nil {
// 		t.Fatalf("Failed to open database: %v", err)
// 	}

// 	// Test closing the database
// 	if err := dbInstance.Close(); err != nil {
// 		t.Fatalf("Failed to close database: %v", err)
// 	}
// }

// func TestGetLastReceiptes(t *testing.T) {
// 	// Create a temporary database file
// 	testDB := "test_receipts.db"
// 	defer os.Remove(testDB) // clean up after the test

// 	// Open the database
// 	dbInstance, err := Open(testDB, "BTC-USD")
// 	if err != nil {
// 		t.Fatalf("Failed to open database: %v", err)
// 	}
// 	defer dbInstance.Close()

// 	// Insert mock receipts into the database
// 	err = dbInstance.Conn().Update(func(tx *bbolt.Tx) error {
// 		bucket := tx.Bucket([]byte("BTC-USD"))
// 		if bucket == nil {
// 			t.Fatal("Bucket not created")
// 		}

// 		// Create dummy receipts
// 		receipts := []db.Reciept{
// 			{ID: "1", Amount: 100.0, Timestamp: time.Now().Unix()},
// 			{ID: "2", Amount: 200.0, Timestamp: time.Now().Unix()},
// 		}

// 		// Insert receipts into the bucket
// 		for _, receipt := range receipts {
// 			data, err := json.Marshal(receipt)
// 			if err != nil {
// 				return err
// 			}
// 			if err := bucket.Put([]byte(receipt.ID), data); err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		t.Fatalf("Failed to insert receipts: %v", err)
// 	}

// 	// Test GetLastReceiptes
// 	receipts, err := dbInstance.GetLastReceiptes(2)
// 	if err != nil {
// 		t.Fatalf("Failed to retrieve receipts: %v", err)
// 	}

// 	// Verify that the correct number of receipts is returned
// 	if len(receipts) != 2 {
// 		t.Fatalf("Expected 2 receipts, got %d", len(receipts))
// 	}

// 	// Check if the order is correct (most recent first)
// 	if receipts[0].ID != "2" {
// 		t.Errorf("Expected first receipt ID '2', got '%s'", receipts[0].ID)
// 	}
// 	if receipts[1].ID != "1" {
// 		t.Errorf("Expected second receipt ID '1', got '%s'", receipts[1].ID)
// 	}
// }

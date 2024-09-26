package postgres

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/EncrypteDL/EDL-TradeBot/internal/db"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var DbB *DB

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
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal(err)
	}

	if err = pool.Client.Ping(); err != nil {
		log.Fatal(err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "16.2-alpine3.19",
		Env: []string{
			"POSTGRES_PASSWORD=pazzword",
			"POSTGRES_USER=uzer",
			"POSTGRES_DB=db",
			"listen_addresses='*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatal(err)
	}

	addr := resource.GetHostPort("5432/tcp")
	dbURL := fmt.Sprintf("postgres://uzer:pazzword@%s/db?sslmode=disable", addr)

	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		DbB, err = Open(dbURL, "BTCUSDT", context.Background())
		return err
	}); err != nil {
		log.Fatal(err)
	}

	code := m.Run()

	if err = pool.Purge(resource); err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func TestInsertReceipt(t *testing.T) {
	err := DbB.InsertReceipt(receipts[1])
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLastReceipts(t *testing.T) {
	err := DbB.InsertReceipt(receipts[0])
	if err != nil {
		t.Fatal(err)
	}

	n := 2

	r, err := DbB.GetLastReceipts(n)
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

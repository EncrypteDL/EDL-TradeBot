package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/EncrypteDL/EDL-TradeBot/internal/db"
	"github.com/redis/go-redis/v9"
)

type DB struct {
	client *redis.Client
	ctx    context.Context
}

var tradingPairSymbol string

func Open(url, symbol string, ctx context.Context) (db *DB, err error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opts)

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	tradingPairSymbol = symbol

	return &DB{client: client, ctx: ctx}, nil
}

func (d *DB) Close() error {
	return d.client.Close()
}

func (d *DB) GetLastReceipts(limit int) ([]*db.Reciept, error) {
	receipts := []*db.Reciept{}

	id, err := d.client.Incr(d.ctx, fmt.Sprintf("id:%s", tradingPairSymbol)).Result()
	if err != nil {
		return nil, err
	}

	for i := id - int64(limit); i < id; i++ {
		key := fmt.Sprintf("receipt:%s:%d", tradingPairSymbol, i)
		val, err := d.client.Get(d.ctx, key).Bytes()
		if err != nil {
			return nil, err
		}

		var receipt db.Reciept
		err = json.Unmarshal(val, &receipt)
		if err != nil {
			return nil, err
		}

		receipts = append(receipts, &receipt)
	}

	return receipts, nil
}

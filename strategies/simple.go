package strategies

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/EncrypteDL/EDL-TradeBot/exchanges/binance"
	"github.com/EncrypteDL/EDL-TradeBot/internal/db/postgres"
	"github.com/EncrypteDL/EDL-TradeBot/internal/db/redis"
	"github.com/EncrypteDL/EDL-TradeBot/internal/db/sqlite"
	"github.com/EncrypteDL/EDL-TradeBot/internal/pricews"
	"github.com/rs/zerolog"
	"go.etcd.io/bbolt"
)

func Run(botCfg *BotCfg) (err error) {
	client := binance.NewClient(botCfg.ServerBaseEndpoint, botCfg.SecretKey, botCfg.ApiKey)

	symbolUpper = strings.ToUpper(botCfg.Base + botCfg.Quote)

	state, err := NewState(botCfg, client)
	if err != nil {
		return
	}

	if botCfg.LogDb {
		switch botCfg.DbBackend {
		case dbBackendSqlite:
			receiptsDb, err = sqlite.Open(botCfg.DbUrl, symbolUpper, context.Background())
		case dbBackendPostgres:
			receiptsDb, err = postgres.Open(botCfg.DbUrl, symbolUpper, context.Background())
		case dbBackendBbolt:
			receiptsDb, err = bbolt.Open(botCfg.DbUrl, symbolUpper)
		case dbBackendRedis:
			receiptsDb, err = redis.Open(botCfg.DbUrl, symbolUpper, context.Background())
		default:
			return fmt.Errorf("invalid db backend: %s", botCfg.DbBackend)
		}

		if err != nil {
			return
		}

		defer receiptsDb.Close()
	}

	logLevel, ok := logLevelMap[botCfg.LogLevel]
	if !ok {
		return fmt.Errorf("invalid log level: %s", botCfg.LogLevel)
	}

	zerolog.SetGlobalLevel(logLevel)

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error)
	priceChan := make(chan pricews.PriceData)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		errChan <- errors.New("signal received, shutting down")
	}()

	go func() {
		errChan <- pricews.Run(botCfg.ServerBaseEndpoint, symbolUpper, priceChan)
	}()

	go func() {
		errChan <- trade(botCfg, state, client, priceChan)
	}()

	return <-errChan
}

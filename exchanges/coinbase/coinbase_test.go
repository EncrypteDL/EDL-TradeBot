package coinbase

import (
	"os"
	"testing"

	"github.com/EncrypteDL/EDL-TradeBot/exchanges"
	"github.com/stretchr/testify/assert"
)

func TestCoinbaseConstructor(t *testing.T) {
	type want struct {
		coinbase *Coinbase
		err      error
	}

	testCases := []struct {
		name  string
		setup func()
		wants want
	}{
		{
			name: "testing with correct env vars",
			setup: func() {
				os.Unsetenv("COINBASE_API_KEY")
				os.Unsetenv("COINBASE_API_SECRET")

				os.Setenv("COINBASE_API_KEY", "FOO")
				os.Setenv("COINBASE_API_SECRET", "BAR")
			},
			wants: want{
				coinbase: &Coinbase{
					APIKey:    "FOO",
					APISecret: "BAR",
				},
			},
		},
		{
			name: "testing with missing api key env var",
			setup: func() {
				os.Unsetenv("COINBASE_API_KEY")
				os.Unsetenv("COINBASE_API_SECRET")

				os.Setenv("COINBASE_API_SECRET", "BAR")
			},
			wants: want{
				err: exchanges.ErrAPIKeyNotSet,
			},
		},
		{
			name: "testing with missing api secret env var",
			setup: func() {
				os.Unsetenv("COINBASE_API_KEY")
				os.Unsetenv("COINBASE_API_SECRET")

				os.Setenv("COINBASE_API_KEY", "FOO")
			},
			wants: want{
				err: exchanges.ErrAPISecretNotSet,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			tt.setup()

			res, err := NewCoinbase()

			assert.Equal(t, tt.wants.coinbase, res, "test: %s", tt.name)
			assert.ErrorIs(t, err, tt.wants.err)
		})
	}
}

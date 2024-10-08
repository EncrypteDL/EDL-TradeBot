package internal

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssetUnit(t *testing.T) {
	testCases := []struct {
		name     string
		asset    Asset
		input    float64
		expected int64
	}{
		{
			name:     "USD units",
			asset:    USD,
			input:    50.01,
			expected: 5001,
		},
		{
			name:     "BTC units",
			asset:    BTC,
			input:    0.01,
			expected: 1000000,
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.asset.Unit(tt.input))
		})
	}
}

func TestAssetFormat(t *testing.T) {
	testCases := []struct {
		name     string
		asset    Asset
		input    int64
		expected string
	}{
		{
			name:     "USD units",
			asset:    USD,
			input:    5001,
			expected: "50.01",
		},
		{
			name:     "BTC units",
			asset:    BTC,
			input:    1000000,
			expected: "0.01",
		},
		{
			name:     "BTC units small",
			asset:    BTC,
			input:    58823,
			expected: "0.00058823",
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.asset.Format(tt.input))
		})
	}
}

func TestAssetUnitStr(t *testing.T) {
	type want struct {
		err error
		res int64
	}

	testCases := []struct {
		name     string
		asset    Asset
		input    string
		expected want
	}{
		{
			name:  "USD units",
			asset: USD,
			input: "50.01",
			expected: want{
				res: 5001,
			},
		},
		{
			name:  "BTC units",
			asset: BTC,
			input: "0.01",
			expected: want{
				res: 1000000,
			},
		},
		{
			name:  "Bad units",
			asset: USD,
			input: "abcde",
			expected: want{
				err: strconv.ErrSyntax,
			},
		},
	}

	for _, tt := range testCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			res, err := tt.asset.UnitStr(tt.input)

			assert.ErrorIs(t, err, tt.expected.err)
			assert.Equal(t, tt.expected.res, res)
		})
	}
}

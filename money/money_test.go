package money

import (
	"testing"
)

func TestMoney_PercentageChange(t *testing.T) {
	tests := map[string]struct {
		inputMoneyStrOriginal string
		inputMoneyStrNew      string
		expected              money
		expectedFmt           string // Fixed typo here from 'epectedFMT'
	}{
		"112.3 => 52%": {
			inputMoneyStrOriginal: "112.35",
			inputMoneyStrNew:      "177.00",
			expected:              money(50000000),
			expectedFmt:           "34.00000000", // Fixed 'epectedFMT' typo
		},
		"2022 => -35.51%": {
			inputMoneyStrOriginal: "2016",
			inputMoneyStrNew:      "1300",
			expected:              money(-3551000000),
			expectedFmt:           "-35.51000000", // Fixed 'epectedFMT' typo
		},
		"1 => 4900%": {
			inputMoneyStrOriginal: "1",
			inputMoneyStrNew:      "50",
			expected:              money(490000000000),
			expectedFmt:           "4900.00000000",
		},
		"0 => 100%": {
			inputMoneyStrOriginal: "0",
			inputMoneyStrNew:      "50",
			expected:              money(10000000000),
			expectedFmt:           "100.00000000",
		},
		"1.56660000 => -85.25%": {
			inputMoneyStrOriginal: "1.56660000",
			inputMoneyStrNew:      "0.231000000",
			expected:              money(-8525000000),
			expectedFmt:           "-85.25000000",
		},
		"0.26660000 => -13.35%": {
			inputMoneyStrOriginal: "0.26660000",
			inputMoneyStrNew:      "0.231000000",
			expected:              money(-1335000000),
			expectedFmt:           "-13.35000000",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			inputMoneyOriginal, err := ParseMoney(tt.inputMoneyStrOriginal)
			if err != nil {
				t.Fatal(err)
			}

			inputMoneyNew, err := ParseMoney(tt.inputMoneyStrNew)
			if err != nil {
				t.Fatal(err)
			}

			got := inputMoneyOriginal.PercentageChange(inputMoneyNew)

			if tt.expected != got {
				t.Fatalf("expected: %d, got: %d", tt.expected, got)
			}

			if tt.expectedFmt != got.FormatMoney(false) {
				t.Fatalf("expected: %s, got: %s", tt.expectedFmt, got.FormatMoney(false))
			}
		})
	}
}

func TestMoney_AmountFromPercentage(t *testing.T) { // Corrected from 'TestBimoney_AmountFromPercentage'
	tests := map[string]struct {
		inputMoneyStr      string
		inputPercentageInt int
		expected           money
		expectedFmt        string
	}{
		"0.21874356 (25%)": {
			inputMoneyStr:      "0.21874356",
			inputPercentageInt: 25,
			expected:           money(5468589),
			expectedFmt:        "0.05468589",
		},
		"112.5 (52%)": {
			inputMoneyStr:      "112.50",
			inputPercentageInt: 52,
			expected:           money(5850000000),
			expectedFmt:        "58.50000000",
		},
		"215.3589 (150%)": {
			inputMoneyStr:      "215.3589",
			inputPercentageInt: 150,
			expected:           money(32303835000),
			expectedFmt:        "323.03835000",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			inputMoney, err := ParseMoney(tt.inputMoneyStr)
			if err != nil {
				t.Fatal(err)
			}

			inputPercentage := money(tt.inputPercentageInt * 100000000) // Fixed issue with percentage conversion

			got := inputMoney.AmountFromPercentage(inputPercentage)

			if tt.expected != got {
				t.Fatalf("expected: %d, got: %d", tt.expected, got)
			}

			if tt.expectedFmt != got.FormatMoney(false) {
				t.Fatalf("expected: %s, got: %s", tt.expectedFmt, got.FormatMoney(false))
			}
		})
	}
}

func TestMoney_PortionOf(t *testing.T) { // Corrected from 'TestBimoney_PortionOf'
	tests := map[string]struct {
		balanceMoneyStr         string
		currentPriceForWholeStr string
		expected                money
		expectedFmt             string
	}{
		"350 as a portion of 0.1501": {
			balanceMoneyStr:         "350",
			currentPriceForWholeStr: "0.15010000",
			expected:                money(233177880000),
			expectedFmt:             "2331.77880000",
		},
		"150 as a portion of 500": {
			balanceMoneyStr:         "150",
			currentPriceForWholeStr: "500",
			expected:                money(30000000),
			expectedFmt:             "0.30000000",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			balanceMoney, err := ParseMoney(tt.balanceMoneyStr)
			if err != nil {
				t.Fatal(err)
			}

			currentPriceForWhole, err := ParseMoney(tt.currentPriceForWholeStr)
			if err != nil {
				t.Fatal(err)
			}

			got := balanceMoney.PortionOf(currentPriceForWhole)

			if tt.expected != got {
				t.Fatalf("expected: %d, got: %d", tt.expected, got)
			}

			if tt.expectedFmt != got.FormatMoney(false) {
				t.Fatalf("expected: %s, got: %s", tt.expectedFmt, got.FormatMoney(false))
			}
		})
	}
}

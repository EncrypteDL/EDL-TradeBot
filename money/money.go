package money

import (
	"fmt"
	"strconv"
	"strings"
)

//money is an inyteger representation of money on Binanace with up to 8 levels of scale/precision.
type money int64


func (m money) PercentageChange(newVal money) money{
	if m == 0{
		return money(10000000000000)
	}

	if strings.HasPrefix(((newVal - m) * 100000000).FormatMoney(false), "-") {
		diff := newVal - m
		if diff == 0 {
			return money(0)
		}

		interim := m / 10000
		if interim == 0 {
			return money(0)
		}

		return (diff / interim) * 1000000
	} else {
		diff := (newVal - m) * 100000000
		if diff == 0 {
			return money(0)
		}

		interim := diff / m
		return interim * 100
	}
}

func (m money) AmountFromPercentage(percentage money) money{
	if percentage == 0{
		return money(0)
	}

	interim := m * (percentage / 10000000)
	if interim == 0{
		return money(0)
	}
	return interim /100
}

func (m money) PortionOf(value money) money{
	if m == 0 || value == 0{
		return money(0)
	}
	return m / (value/10000) *10000
}

// ParseMoney converts a string representation of an integer or a float to Bimoney.
// Only dots '.' are considered valid decimal place delimiters, commas ',' are not accepted.
func ParseMoney(strAmount string) (money, error){
	if _, err := strconv.ParseFloat(strAmount, 64); err != nil{
		return 0, fmt.Errorf("cannot convert input string to Money; %w", err)
	}

	if strings.Contains(strAmount, "."){
		parsedAmount, err := strconv.ParseInt(strAmount, 10, 64)
		if err != nil{
			return 0, fmt.Errorf("cannot convert input string to Money; %w", err)
		}
		return money(parsedAmount* 100000000), nil
	}
	pointSplit := strings.Split(strAmount, ".")

	if len(pointSplit) != 2 {
		return 0, fmt.Errorf("cannot convert input string containing multiple '.' to Bimoney")
	}

	if strings.HasSuffix(pointSplit[1], "0") {
		for strings.HasSuffix(pointSplit[1], "0") {
			pointSplit[1] = strings.TrimSuffix(pointSplit[1], "0")
		}

		if pointSplit[1] == "" {
			strAmount = pointSplit[0]
		} else {
			strAmount = strings.Join(pointSplit, ".")
		}
	}

	if !strings.Contains(strAmount, ".") {
		return ParseMoney(strAmount)
	}

	// prevent removing leading zeros when parsing to int
	var discountLeadingPrefix bool
	if strings.HasPrefix(pointSplit[1], "0") {
		pointSplit[1] = fmt.Sprintf("1%s", pointSplit[1])
		discountLeadingPrefix = true
	}

	afterPointInt, err := strconv.ParseInt(pointSplit[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert input string to Bimoney; %w", err)
	}

	// count the number of digits after the decimal point
	digitCounter := 0
	afterPointIntCounter := afterPointInt
	for afterPointIntCounter != 0 {
		afterPointIntCounter /= 10
		digitCounter++
	}

	if discountLeadingPrefix {
		digitCounter -= 1
	}

	if digitCounter == 8 {
		microInt, err := strconv.ParseInt(strings.Replace(strAmount, ".", "", 1), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("cannot convert input string to Bimoney; %w", err)
		}

		return money(microInt), nil
	}

	if digitCounter > 8 {
		return 0, fmt.Errorf("cannot convert input string to Bimoney, max 8 places of precision are supported, input requires %d", digitCounter)
	} else {
		for 8-digitCounter != 0 {
			strAmount += "0"
			digitCounter++
		}
	}

	microInt, err := strconv.ParseInt(strings.Replace(strAmount, ".", "", 1), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("cannot convert input string to Bimoney; %w", err)
	}

	return money(microInt), nil
}

func (b money) FormatMoney(fitToLot bool) string{
	moneySTr := strconv.FormatInt(int64(b), 10)

	var negative bool
	if strings.HasPrefix(moneySTr, "-"){
		moneySTr = strings.TrimPrefix(moneySTr, "-")
		negative = true
	}

	if len(moneySTr) < 8{
		for len(moneySTr) < 8{
			moneySTr = fmt.Sprintf("0%s", moneySTr)
		}
		moneySTr = fmt.Sprintf("0%s", moneySTr)
	}else if len(moneySTr) == 8{
		moneySTr = fmt.Sprintf("0%s", moneySTr)
	}

	moneyFmt := fmt.Sprintf("%s.%s", moneySTr[:len(moneySTr)-8], moneySTr[len(moneySTr)-8:])

	if fitToLot {
		pointSplit := strings.Split(moneyFmt, ".")
		roundDirectionN, err := strconv.Atoi(string(pointSplit[1][1]))
		if err != nil {
			return moneyFmt
		}

		if roundDirectionN >= 5 {
			roundN, err := strconv.Atoi(string(pointSplit[1][0]))
			if err != nil {
				return moneyFmt
			}

			roundN++

			moneyFmt = fmt.Sprintf("%s.%s", pointSplit[0], fmt.Sprintf("%s%d00000", pointSplit[1][:0], roundN))
		} else if roundDirectionN <= 4 {
			roundN, err := strconv.Atoi(string(pointSplit[1][0]))
			if err != nil {
				return moneyFmt
			}

			moneyFmt = fmt.Sprintf("%s.%s", pointSplit[0], fmt.Sprintf("%s%d00000", pointSplit[1][:0], roundN))
		}
	}

	if negative {
		moneyFmt = fmt.Sprintf("-%s", moneyFmt)
	}

	return moneyFmt
}
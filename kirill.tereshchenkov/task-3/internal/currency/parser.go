package currency

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

var errInvalidCurrencyData = errors.New("invalid currency data")

type xmlDocument struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []xmlValute `xml:"Valute"`
}

type xmlValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

// ParseXML читает XML-файл и преобразует его в срез валют.
func ParseXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var document xmlDocument

	err = decoder.Decode(&document)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	currencies := make([]Currency, 0, len(document.Valutes))

	for i, valute := range document.Valutes {
		charCode := strings.TrimSpace(valute.CharCode)

		if valute.NumCode <= 0 || charCode == "" {
			return nil, fmt.Errorf("%w at index %d", errInvalidCurrencyData, i)
		}

		rawValue := strings.ReplaceAll(
			strings.TrimSpace(valute.Value),
			",",
			".",
		)

		value, err := strconv.ParseFloat(rawValue, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"failed to parse value for %s: %w",
				charCode,
				err,
			)
		}

		parsedCurrency := Currency{
			NumCode:  valute.NumCode,
			CharCode: charCode,
			Value:    value,
		}
		currencies = append(currencies, parsedCurrency)
	}

	return currencies, nil
}

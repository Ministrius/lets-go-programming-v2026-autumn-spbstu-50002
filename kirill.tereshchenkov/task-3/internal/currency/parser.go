package currency

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type xmlDocument struct {
	XMLName xml.Name    `xml:"ValCurs"`
	Valutes []xmlValute `xml:"Valute"`
}
type xmlValute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ParseXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read XML: %w", err)
	}

	var document xmlDocument

	err = xml.Unmarshal(data, &document)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XML: %w", err)
	}

	currencies := make([]Currency, 0, len(document.Valutes))

	for i, valute := range document.Valutes {
		charCode := strings.TrimSpace(valute.CharCode)

		if valute.NumCode <= 0 || charCode == "" {
			return nil, fmt.Errorf("invalid currency data at index %d", i)
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
		currencies = append(currencies, Currency{
			NumCode:  valute.NumCode,
			CharCode: charCode,
			Value:    value,
		})
	}

	return currencies, nil
}

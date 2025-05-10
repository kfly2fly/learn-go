package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))
	for i, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return nil, errors.New("Failed to convert string to float")
		}

		floats[i] = floatVal
	}
	return floats, nil
}

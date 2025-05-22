package Country

import (
	"github.com/thereisnoplanb/globalization/Continent"
	"github.com/thereisnoplanb/globalization/CountryCode"
)

type Country struct {
	Name            string
	Code            CountryCode.Enum
	Continent       Continent.Enum
	PostalCodeRegex *string
}

func getPtr[T any](value T) *T {
	return &value
}

var Poland Country = Country{
	Name:            "Poland",
	Code:            CountryCode.Poland,
	Continent:       Continent.Europe,
	PostalCodeRegex: getPtr("[0-9]{2}-[0-9]{3}"),
}

var Germany Country = Country{
	Name:            "Germany",
	Code:            CountryCode.Germany,
	Continent:       Continent.Europe,
	PostalCodeRegex: getPtr("[0-9]{5}"),
}

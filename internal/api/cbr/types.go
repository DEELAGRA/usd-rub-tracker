package cbr

// "USD":{"ID":"R01235","NumCode":"840","CharCode":"USD","Nominal":1,"Name":"Доллар США","Value":75.2246,"Previous":75.3448}
type CBRResponse struct {
	Date   string              `json:"Date"`
	Valute map[string]Currency `json:"Valute"`
}

type Currency struct {
	ID       string  `json:"ID"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}
